package core_test

import (
	"Api/src/core"
	"errors"
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTratarErro(t *testing.T){
	t.Parallel()

	t.Run("Deve_Lancar_Erro", func(t *testing.T){
		// arrange
		var resultadoChamada interface{}
		mensagemErro := "TESTANDO"
		erroEsperado := errors.New(mensagemErro)
		logFatalExecutado := false

		// act
		monkey.Patch(log.Fatal, func(v ...interface{}){
			logFatalExecutado = true
			resultadoChamada = v[0]
		})
		core.TratarErro(errors.New(mensagemErro))

		// assert
		defer func(){
			assert.True(t, logFatalExecutado)
			assert.Equal(t, resultadoChamada, erroEsperado)
		}()
	})

	t.Run("Nao_Deve_Lancar_Erro_Quando_For_Nil", func(t *testing.T){
		//assert

		// act
		core.TratarErro(nil)

		// assert
		assert.Equal(t, true, true)
	})
}
