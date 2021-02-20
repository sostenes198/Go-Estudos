package rotas_base

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	Nome               string
	RequerAutenticacao bool
}

func ConfigurarRotas(router *mux.Router, rotasControllers []IRota) *mux.Router {
	for _, rotasController := range rotasControllers {
		for _, rota := range rotasController.ListarRotas() {
			router.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo).Name(rota.Nome)
		}
	}
	return router
}
