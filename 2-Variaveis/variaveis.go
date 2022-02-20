package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string
		variavel4 string
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Contante1"
	fmt.Println(constante1)

	const constante2 = "Contante1"
	fmt.Println(constante2)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
