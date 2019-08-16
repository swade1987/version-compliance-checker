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

	// mapping endpoints
	endpoints := compliant.Endpoints{
		GetEndpoint:      compliant.MakeGetEndpoint(srv),
		StatusEndpoint:   compliant.MakeStatusEndpoint(srv),
		ValidateEndpoint: compliant.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("version-compliance-checker is listening on port:", *httpAddr)
		handler := compliant.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
