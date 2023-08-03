package enderecos_test

import (
	"introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Abc", "Rua"},
		{"Avenida Abc", "Avenida"},
		{"Estrada Abc", "Estrada"},
		{"Rodovia Abc", "Rodovia"},
		{"Praca Abc", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoRecebido := enderecos.TipoDeEndereco(cenario.enderecoInserido)
		if tipoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e receu %s", cenario.retornoEsperado, tipoRecebido)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
