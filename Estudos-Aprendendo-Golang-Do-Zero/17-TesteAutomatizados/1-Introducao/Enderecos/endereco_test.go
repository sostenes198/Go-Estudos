// TESTE DE UNIDADE
package Enderecos_test

import (
	. "Introducao/Enderecos"
	"testing"
)

type cenariosDeTeste struct{
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T){
	t.Parallel()

	cenariosTeste := []cenariosDeTeste{
		{"Rua Andradas", "Rua"},
		{"Avenida Andradas", "Avenida"},
		{"Estrada Andradas", "Estrada"},
		{"Rodovia Andradas", "Rodovia"},
		{"Praça Andradas", "Tipo Inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA DOS BOBOS", "Avenida"},
		{"ESTRADA DOS BOBOS", "Estrada"},
		{"RODOVIA DOS BOBOS", "Rodovia"},
	}


	for _, cenario := range cenariosTeste{
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if cenario.retornoEsperado != tipoEnderecoRecebido {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}


