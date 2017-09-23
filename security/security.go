package security

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// Config for the server
type Config struct {
	secure *secure.Secure
}

// New create the new server security configuration
func New() *Config {
	return &Config{
		secure: secure.New(secure.Options{
			FrameDeny:          true,
			BrowserXssFilter:   true,
			ContentTypeNosniff: true,
		}),
	}
}

// Apply the current configuration to the server
func (s *Config) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.secure.Process(c.Writer, c.Request)
		// If an error occurs do not continue
		if err != nil {
			c.Abort()
			return
		}
		// If the response is a redirection, avoid header rewrite
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
	}
}
