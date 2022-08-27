package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/taidevops/dashboard/src/app/backend/client"
	"github.com/spf13/pflag"
	"github.com/taidevops/dashboard/src/app/backend/args"
	"github.com/taidevops/dashboard/src/app/backend/handler"
)

var (
	argInsecurePort = pflag.Int("insecure-port", 9090, "port to listen to for incoming HTTP requests")
)

func main() {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = flag.CommandLine.Parse(make([]string, 0)) // Init for glog calls in kubernetes packages

	// Initializes dashboard arguments holders
	initArgHolder()

	clientManager := client.NewClientManager()
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

func initArgHolder() {
	builder := args.GetHolderBuilder()
}

func handleFatalInitError(err error) {
	log.Fatalf("Error while initializing connection to Kubernetes apiserver. "+
		"This most likely means that the cluster is misconfigured "+
		"%s\n", err)
}
