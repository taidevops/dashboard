package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/taidevops/dashboard/src/app/backend/handler"
)

func main() {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)

	apiHandler, err := handler.CreateHTTPAPIHandler()

	if err != nil {
		handleFatalInitError(err)
	}

	healthHandler := 

	// Run a HTTP server that serves static public files from './public' and handles API calls.
	http.Handle("/", handler.AppHandler())
	http.Handle("/api/", apiHandler)

	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 8080)
	go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()
	select {}
}

func handleFatalInitError(err error) {
	log.Fatalf("Error while initializing connection to Kubernetes apiserver. "+
		"This most likely means that the cluster is misconfigured "+
		"%s\n", err)
}
