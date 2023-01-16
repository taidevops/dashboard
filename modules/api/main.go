package main

import (
	"crypto/elliptic"
	"log"
	"os"
)

var (

)

func main() {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = flag.CommandLine.Parse(make([]string, 0))

	initArgHolder()
}