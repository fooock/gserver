package controller

import (
	"github.com/fooock/gserver/auth"
	"github.com/fooock/gserver/common"
	"github.com/fooock/gserver/user"
	"github.com/gin-gonic/gin"
)

func userEndpoints(router *gin.Engine) {
	for _, value := range user.UserRoutes {
		// Create one group for each endpoint
		userV1 := router.Group(common.APIV1)
		// Post method for user don't need auth because is the
		// register or sign in proces...
		if value.Method == common.POST {
			userV1.POST(value.Pattern, value.Handler)
		}
		userV1.Use(auth.Auth())
		{
			if value.Method == common.GET {
				userV1.GET(value.Pattern, value.Handler)
			}
		}
	}
}
