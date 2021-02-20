package main

import "fmt"

func main()  {
	fmt.Println("Estrutuas de controle")

	numero := 10

	if numero > 15{
		fmt.Println("Maior que 15")
	} else if numero < 5 {
		fmt.Println("Menor que 5")
	} else{
		fmt.Println("Sei lá")
	}

	// If Init
	if outroNumero := numero; outroNumero > 0{
		fmt.Println(outroNumero)
	}
}
