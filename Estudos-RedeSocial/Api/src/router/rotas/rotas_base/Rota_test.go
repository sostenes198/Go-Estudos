package rotas_base_test

import (
	. "Api/src/router/rotas/rotas_base"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type MockedRotas struct {
	mock.Mock
}

func (m *MockedRotas) ListarRotas() []Rota{
	var result = m.Called()
	return result.Get(0).([]Rota)
}

func TestConfigurarRotas(t *testing.T) {
	t.Parallel()

	t.Run("Deve_Configurar_Rotas", func(t *testing.T) {
		// arrange
		router := mux.NewRouter()
		mockRotas := new(MockedRotas)

		nomeRotaPost := "POST_TESTE"
		nomeRotaGet := "GET_TESTE"

		rotaPost := Rota{
			Uri:                "/teste",
			Metodo:             http.MethodPost,
			Funcao: func(w http.ResponseWriter, r *http.Request) {

			},
			Nome:               nomeRotaPost,
			RequerAutenticacao: false,
		}
		rotaGet := Rota{
			Uri:                "/teste",
			Metodo:             http.MethodGet,
			Funcao: func(w http.ResponseWriter, r *http.Request) {

			},
			Nome:               nomeRotaGet,
			RequerAutenticacao: false,
		}

		mockRotas.On("ListarRotas").Return([]Rota{
			rotaPost,
			rotaGet,
		})

		// act
		var resultadoRouter = ConfigurarRotas(router, []IRota{mockRotas})
		resultadoRotaPostTest := resultadoRouter.GetRoute(nomeRotaPost)
		resultadoRotaGetTest := resultadoRouter.GetRoute(nomeRotaGet)



		// assert
		assert.NotNil(t, resultadoRouter)
		assert.NotNil(t, resultadoRotaPostTest)
		assert.NotNil(t, resultadoRotaGetTest)
	})
}
