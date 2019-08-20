package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	"github.com/swade1987/version-compliance-checker/pkg/devicemapping"
	"github.com/swade1987/version-compliance-checker/pkg/devicevalidation"
)

func main() {

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var dmRepo = devicemapping.NewInMemRepository()

	iosRequiredVersion, err := os.LookupEnv("IOS_REQUIRED_VERSION")
	if err == false {
		logger.Log("err", "Required environment variable IOS_REQUIRED_VERSION missing.")
		os.Exit(-1)
	}
	if err := dmRepo.AddRequiredVersion("ios", iosRequiredVersion); err != nil {
		logger.Log("err", err)
		os.Exit(-1)
	}

	androidRequiredVersion, err := os.LookupEnv("ANDROID_REQUIRED_VERSION")
	if err == false {
		logger.Log("err", "Required environment variable ANDROID_REQUIRED_VERSION missing.")
		os.Exit(-1)
	}
	if err := dmRepo.AddRequiredVersion("android", androidRequiredVersion); err != nil {
		logger.Log("err", err)
		os.Exit(-1)
	}

	logger.Log("iosRequiredVersion", iosRequiredVersion)
	logger.Log("androidRequiredVersion", androidRequiredVersion)

	fieldKeys := []string{"method"}

	var dvs devicevalidation.Service
	dvs = devicevalidation.NewService(dmRepo)
	dvs = devicevalidation.NewLoggingService(logger, dvs)
	dvs = devicevalidation.NewInstrumentingService(kitprometheus.NewCounterFrom(
		stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "device_validation",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "device_validation",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		dvs)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/devicevalidation/v1", devicevalidation.MakeHandler(dvs, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
