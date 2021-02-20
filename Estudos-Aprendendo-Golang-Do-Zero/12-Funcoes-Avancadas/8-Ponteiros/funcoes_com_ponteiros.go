package main

import "fmt"

func inverterSinal(numero int) int{
	return numero * -1
}

func inverterSinalComPonteiro(numero *int){
	*numero = *numero * -1
}

func main(){
	fmt.Println("Funções com ponteiros")

	numero := 10
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
