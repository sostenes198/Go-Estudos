package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Tipos de dados")

	// NUMEROS INTEIROS
	numero := 1000000
	fmt.Println(numero)

	// uint é um inteiro sem sinal
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	// NUMEROS INTEIROS

	// NÚMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	numeroReal2 := 4123.321
	fmt.Println(numeroReal2)
	// FIM NÚMEROS REAIS

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	// FIM STRINGS

	var boolean bool = true // valor default é false
	fmt.Println(boolean)

	var erro error = errors.New("Erro interno.")
	fmt.Println(erro)
}
