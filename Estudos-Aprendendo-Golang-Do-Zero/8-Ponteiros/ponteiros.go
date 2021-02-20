package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

func main(){
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int = 100
	var ponteiro *int = &variavel3

	// & - para acessar o endereço de memória
	// * - Para acessar o valor armazenado no endereço de memória do ponteiro
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 1
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	*ponteiro = 1010
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	p1 := pessoa{"soso", 25}
	p2 := p1

	fmt.Println(p1, p2)

	p1.nome = "Jose"
	p1.idade = 40
	fmt.Println(p1, p2)
}
