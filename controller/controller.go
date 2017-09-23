package controller

import (
	"fmt"

	"github.com/fooock/gserver/common"
	"github.com/fooock/gserver/security"
	"github.com/gin-gonic/gin"
)

// Server is the program instance to un all the things
type Server struct {
	router *gin.Engine
}

// New create a new server
// When the server is created, if the release flag is present the
// security configuration is applied
func New() *Server {
	customRouter := gin.Default()
	// If the server is in release mode we apply the security configuration
	if gin.Mode() == "release" {
		sec := security.New()
		customRouter.Use(sec.Apply())
	}
	// register user resources
	userEndpoints(customRouter)

	server := &Server{
		router: customRouter,
	}
	return server
}

// Start is the function to initialize the server in the given host:port
func (s *Server) Start(options common.Options) {
	s.router.Run(fmt.Sprintf("%v:%v", options.Host, options.Port))
}
