package main

import (
	"errors"
	"fmt"
)

func main(){
	// int8, int16, int32, int64 - Número inteiros
	// uint unsygned - Int sem sinal
	// Alias para int32 = rune
	// Alias para int8 = byte

	// float34, float64 - Números flutuantes

	// string - Caracteres

	// tipo 0 (Todas as variáveis são sempre iniciadas, com valor default)

	// bool (true, false)

	// error (Valor zero dele é <nil>

	var error error = errors.New("Erro interno")

	fmt.Println(error)
}
