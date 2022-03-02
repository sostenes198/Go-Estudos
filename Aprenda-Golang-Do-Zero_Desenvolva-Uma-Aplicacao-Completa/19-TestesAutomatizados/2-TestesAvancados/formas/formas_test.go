package formas_test

import (
	"math"
	. "testes-avancados/formas"
	"testing"
	"time"
)

func TestArea(t *testing.T) {
	t.Parallel()
	t.Run("Retângulo", func(t *testing.T) {
		t.Parallel()
		retangulo := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
		time.Sleep(time.Millisecond * 1000)
	})

	t.Run("Círculo", func(t *testing.T) {
		t.Parallel()
		circulo := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
		time.Sleep(time.Millisecond * 1000)
	})

	time.Sleep(time.Millisecond * 500)
}

type cenarioTesteArea struct {
	forma        Forma
	areaEsperada float64
}

func TestAreaComCenario(t *testing.T) {
	t.Parallel()

	cenariosDeTest := []cenarioTesteArea{
		{Retangulo{Altura: 10, Largura: 12}, 120},
		{Circulo{Raio: 10}, math.Pi * 100},
	}

	for _, cenario := range cenariosDeTest {
		areaRecebida := cenario.forma.Area()
		if areaRecebida != cenario.areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, cenario.areaEsperada)
		}
	}

	time.Sleep(time.Millisecond * 2000)
}
