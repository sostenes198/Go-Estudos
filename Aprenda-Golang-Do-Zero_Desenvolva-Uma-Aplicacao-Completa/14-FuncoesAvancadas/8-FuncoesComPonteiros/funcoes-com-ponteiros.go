package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Funções com Ponteiros")
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	fmt.Println("----")

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
