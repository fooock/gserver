package auth

import (
	"github.com/gin-gonic/gin"
)

// Auth is a handler function to require authentication to endpoints
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Write([]byte("auth required for endpoint"))
	}
}
