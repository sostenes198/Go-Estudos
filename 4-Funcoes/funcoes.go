package main

import "fmt"

func main() {
	fmt.Println("Funções")

	soma := somar(1, 3)
	fmt.Println(soma)

	fmt.Println(calculosMatematicos(3, 1))

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}

func somar(n1 int, n2 int) int {
	f := func(texto string) {
		fmt.Println(texto)
	}
	f("Texto da função 1.")

	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}
