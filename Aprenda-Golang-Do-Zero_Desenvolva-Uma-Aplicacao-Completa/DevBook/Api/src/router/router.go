package router

import (
	"devbook/src/middlewares"
	"devbook/src/router/routers"
	"github.com/gorilla/mux"
)

// Generate vai retornar um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()

	allRouters := routers.UserRouters
	allRouters = append(allRouters, routers.LoginRouters)

	for _, route := range allRouters {

		if route.RequiredAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)
		}else{
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
