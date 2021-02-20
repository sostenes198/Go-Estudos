package main

import "fmt"

func soma(numeros ...int) (total int){
	for _, numero := range numeros{
		total +=numero
	}
	return
}

func escrever(texto string, numeros ...int){
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variáticas")
	fmt.Println(soma(1,2,3,4,5))
	escrever("Olá mundo", 1,2,3,4,5)
}
