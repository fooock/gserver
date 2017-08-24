package controller

import (
	"github.com/fooock/gserver/common"
	"github.com/fooock/gserver/user"
	"github.com/gin-gonic/gin"
)

func userEndpoints(router *gin.Engine) {
	for _, value := range user.UserRoutes {
		userV1 := router.Group(common.APIV1)
		{
			if value.Method == common.GET {
				userV1.GET(value.Pattern, value.Handler)
			}
			if value.Method == common.POST {
				userV1.POST(value.Pattern, value.Handler)
			}
		}
	}
}
