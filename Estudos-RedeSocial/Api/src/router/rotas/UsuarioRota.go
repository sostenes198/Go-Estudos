package rotas

import (
	"Api/src/controllers"
	"Api/src/router/rotas/rotas_base"
	"net/http"
)

type UsuarioRota struct{}

func (_ UsuarioRota) ListarRotas() []rotas_base.Rota {
	var usuarioController = controllers.UsuarioController{}
	return []rotas_base.Rota{
		{
			Uri:                "/usuarios",
			Metodo:             http.MethodPost,
			Funcao:             usuarioController.Criar,
			Nome: "POST_USUARIO",
			RequerAutenticacao: false,
		},
		{
			Uri:                "/usuarios",
			Metodo:             http.MethodGet,
			Funcao:             usuarioController.Listar,
			Nome: "GET_USUARIO",
			RequerAutenticacao: false,
		},
		{
			Uri:                "/usuarios/{id}",
			Metodo:             http.MethodGet,
			Funcao:             usuarioController.ListarPorId,
			Nome: "GET_BY_ID_USUARIO",
			RequerAutenticacao: false,
		},
		{
			Uri:                "/usuarios/{id}",
			Metodo:             http.MethodPut,
			Funcao:             usuarioController.Atualizar,
			Nome: "PUT_USUARIO",
			RequerAutenticacao: false,
		},
		{
			Uri:                "/usuarios/{id}",
			Metodo:             http.MethodDelete,
			Funcao:             usuarioController.Excluir,
			Nome: "DELETE_USUARIO",
			RequerAutenticacao: false,
		},
	}
}
