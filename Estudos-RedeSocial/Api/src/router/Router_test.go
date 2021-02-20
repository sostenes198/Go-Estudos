package router_test

import (
	"Api/src/router"
	"Api/src/router/rotas/rotas_base"
	"github.com/bouk/monkey"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGerar(t *testing.T){
	t.Parallel()

	t.Run("Deve_Gerar_Rotas", func(t *testing.T){
		// arrange
		resultadoEsperado := mux.NewRouter()
		configuracoesRotasExecutada := false

		monkey.Patch(rotas_base.ConfigurarRotas, func(router *mux.Router, rotasControllers []rotas_base.IRota) *mux.Router{
			configuracoesRotasExecutada = true
			return mux.NewRouter()
		})

		// act
		resultado := router.Gerar()

		//assert
		defer func(){
			assert.Equal(t, resultado, resultadoEsperado)
			assert.True(t, configuracoesRotasExecutada)
		}()

	})
}

