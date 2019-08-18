package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/swade1987/version-compliance-checker/pkg/compliant"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()
	ctx := context.Background()

	srv := compliant.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	iosRequiredVersion, err := os.LookupEnv("IOS_REQUIRED_VERSION")
	if err == false {
		log.Fatal("Required environment variable IOS_REQUIRED_VERSION missing.")
	}

	androidRequiredVersion, err := os.LookupEnv("ANDROID_REQUIRED_VERSION")
	if err == false {
		log.Fatal("Required environment variable ANDROID_REQUIRED_VERSION missing.")
	}

	log.Println("IOS_REQUIRED_VERSION: ", iosRequiredVersion)
	log.Println("ANDROID_REQUIRED_VERSION: ", androidRequiredVersion)

	// mapping endpoints
	endpoints := compliant.Endpoints{
		StatusEndpoint:   compliant.MakeStatusEndpoint(srv),
		ValidateEndpoint: compliant.MakeValidateEndpoint(srv, iosRequiredVersion, androidRequiredVersion),
	}

	// HTTP transport
	go func() {
		log.Println("version-compliance-checker is listening on port:", *httpAddr)
		handler := compliant.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
