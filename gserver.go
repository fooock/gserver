package main

import (
	"flag"
	"fmt"

	"github.com/fooock/gserver/common"
	"github.com/fooock/gserver/controller"
)

const version = "v0.1"

func main() {
	fmt.Printf("GServer %v\n", version)

	// Get the command line options
	port := flag.String("p", "8080", "Port where the server will be listening")
	host := flag.String("x", "0.0.0.0", "Host to attach the server")
	flag.Parse()

	// Create the options for the server
	options := &common.Options{
		Port: *port,
		Host: *host,
	}
	// Create and start the server with the given options
	server := controller.New()
	server.Start(*options)
}
