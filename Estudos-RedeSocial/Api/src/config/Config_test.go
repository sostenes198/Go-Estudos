package config_test

import (
	. "Api/src/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCarregar(t *testing.T) {
	t.Parallel()

	t.Run("Deve_Carregar_Variaveis_De_Ambiente", func(t *testing.T) {
		// arrange
		resultadoStringConexaoEsperado := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
		resultadoPortaEsperado := 5000

		os.Setenv(ApiPort, "asd")

		// act
		Carregar()

		// assert
		assert.Equal(t, StringConexao, resultadoStringConexaoEsperado)
		assert.Equal(t, resultadoPortaEsperado, resultadoPortaEsperado)

	})
}
