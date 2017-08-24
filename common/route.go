package common

import "github.com/gin-gonic/gin"

const (
	// APIV1 is the base for all server endpoints
	APIV1 = "api/v1"
	// GET http method
	GET = "GET"
	// POST http method
	POST = "POST"
)

// Route is the struc to create rest resources
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler func(c *gin.Context)
}

// Routes is a list of Route
type Routes []Route
