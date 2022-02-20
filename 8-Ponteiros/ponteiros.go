package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	variavel1 := 10
	variavel2 := 10

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)
	fmt.Println()

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // desreferenciação

	fmt.Println()

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	fmt.Println()

	*ponteiro = 10

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
