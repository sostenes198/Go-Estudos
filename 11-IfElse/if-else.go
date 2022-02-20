package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 0

	if numero >= 15 {
		fmt.Println("Maior ou igual a 15")
	} else {
		fmt.Println("Menor que  15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Outro número é maior que 0")
	}

}
