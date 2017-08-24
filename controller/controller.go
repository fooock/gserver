package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server is the program instance to un all the things
type Server struct {
	router *gin.Engine
}

// New create a new server
func New() *Server {
	customRouter := gin.Default()
	// register user resources
	userEndpoints(customRouter)

	server := &Server{
		router: customRouter,
	}
	return server
}

// Start is the function to initialize the server in the given port
func (s *Server) Start(port string) {
	s.router.Run(fmt.Sprintf(":%v", port))
}
