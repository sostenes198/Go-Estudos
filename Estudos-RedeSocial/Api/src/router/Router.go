package router

import (
	"Api/src/router/rotas"
	"Api/src/router/rotas/rotas_base"
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	muxRouter := mux.NewRouter()

	return rotas_base.ConfigurarRotas(muxRouter,
		[]rotas_base.IRota{
			rotas.UsuarioRota{},
		})
}
