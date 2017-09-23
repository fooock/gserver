package user

import "github.com/fooock/gserver/common"

// UserRoutes store all routes for the User resource
var UserRoutes = common.Routes{
	common.Route{
		Name:    "Get all users",
		Method:  "GET",
		Pattern: "/users",
		Handler: All,
	},
	common.Route{
		Name:    "Get user by id",
		Method:  "GET",
		Pattern: "/user/:id",
		Handler: ByID,
	},
	common.Route{
		Name:    "Register a new user",
		Method:  "POST",
		Pattern: "/user/register",
		Handler: Register,
	},
	common.Route{
		Name: "Sign in",
		Method: "POST",
		Pattern: "user/login",
		Handler: Login,
	},
}
