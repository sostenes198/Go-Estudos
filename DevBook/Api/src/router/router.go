package router

import (
	"devbook/src/router/routers"
	"github.com/gorilla/mux"
)

// Generate vai retornar um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()

	routes := routers.UserRouters

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
