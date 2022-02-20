package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, valor := range numeros {
		total += valor
	}

	return total
}

func main() {
	fmt.Println("Funções Variáticas")

	fmt.Println(soma(1, 2, 3, 4, 5, 6))
}
