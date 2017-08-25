package main

import (
	"flag"
	"fmt"

	"github.com/fooock/gserver/common"
	"github.com/fooock/gserver/controller"
	"github.com/gin-gonic/gin"
)

const (
	version = "v0.1.0"
	header  = `
 ________  _________                                
/  _____/ /   _____/ ______________  __ ___________ 
/   \  ___ \_____  \_/ __ \_  __ \  \/ // __ \_  __ \
\    \_\  \/        \  ___/|  | \/\   /\  ___/|  | \/
 \______  /_______  /\___  >__|    \_/  \___  >__|   
        \/        \/     \/                 \/     ` + version
)

func main() {
	fmt.Println(header)

	// Get the command line options
	port := flag.String("p", "8080", "Port where the server will be listening")
	host := flag.String("x", "0.0.0.0", "Host to attach the server")
	// Option to start the server in release mode (the default option) or in debug mode
	// using the -dev option
	debug := flag.Bool("dev", false, "Initialize the server is debug mode")
	flag.Parse()

	// Create the options for the server
	options := &common.Options{
		Port: *port,
		Host: *host,
	}

	// if the debug mode is not enabled, start the server in release mode!
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create and start the server with the given options
	server := controller.New()
	server.Start(*options)
}
