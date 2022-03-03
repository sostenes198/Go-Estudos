package routers

import (
	"devbook/src/controllers"
	"net/http"
)

var UserRouters = []router{
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   controllers.List,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Func:                   controllers.GetById,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   controllers.Create,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Func:                   controllers.Update,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Func:                   controllers.Delete,
		RequiredAuthentication: true,
	},
}
