package Formas_test

import (
	. "2-Testes-Avancados/Formas"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Parallel()
	t.Run("Retângulo", func(t *testing.T) {
		ret :=  Retangulo{Altura: 10, Largural: 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()
		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da espera %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		cir := Circulo{Raio: 10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.Area()
		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da espera %f", areaRecebida, areaEsperada)
		}
	})
}
