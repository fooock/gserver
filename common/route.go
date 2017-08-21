package common

import "github.com/gin-gonic/gin"

// Route is the struc to create rest resources
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler func(c *gin.Context)
}

// Routes is a list of Route
type Routes []Route
