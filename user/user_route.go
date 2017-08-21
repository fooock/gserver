package user

import "github.com/fooock/gserver/common"

// UserRoutes store all routes for the User resource
var UserRoutes = common.Routes{
	common.Route{
		Name:    "Get all users",
		Method:  "GET",
		Pattern: "/users",
		Handler: AllUsers,
	},
	common.Route{
		Name:    "Get user by id",
		Method:  "GET",
		Pattern: "/user/:id",
		Handler: UserById,
	},
}
