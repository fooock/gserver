package main

import (
	"flag"
	"fmt"

	"github.com/fooock/gserver/controller"
)

const version = "v0.1"

var port string

func main() {
	fmt.Printf("GServer %v\n", version)

	// Get the command line options
	flag.StringVar(&port, "p", "8080", "Port where the server will be listening")
	flag.Parse()

	server := controller.New()
	server.Start(port)
}
