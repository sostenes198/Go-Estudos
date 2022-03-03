package routers

import (
	controllers "devbook/src/controllers"
	"net/http"
)

var UserRouters = []router{
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   controllers.List,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Func:                   controllers.GetById,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   controllers.Create,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Func:                   controllers.Update,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Func:                   controllers.Delete,
		RequiredAuthentication: false,
	},
}
