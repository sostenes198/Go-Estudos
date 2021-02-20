package rotas_test

import (
	"Api/src/controllers"
	. "Api/src/router/rotas"
	"Api/src/router/rotas/rotas_base"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListarRotas(t *testing.T) {
	t.Parallel()

	t.Run("Deve_Listar_Rotas_Usuario", func(t *testing.T) {
		// arrange
		usuarioRota := UsuarioRota{}
		var usuarioController = controllers.UsuarioController{}

		resultadoRotasEsperado := []rotas_base.Rota{
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

		//act
		rotas := usuarioRota.ListarRotas()


		// assert
		assert.True(t, RotasArrayEqual(rotas, resultadoRotasEsperado))
	})
}

func RotasArrayEqual(rotas []rotas_base.Rota, rotasEsperadas []rotas_base.Rota) bool {
	if len(rotas) != len(rotasEsperadas) {
		return false
	}
	for indice, rota := range rotas {
		if rota.Nome != rotasEsperadas[indice].Nome {
			return false
		}
	}
	return true
}
