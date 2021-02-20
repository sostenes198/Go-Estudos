package main

import "fmt"

func main(){
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constant1 string = "Constante1"
	fmt.Println(constant1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
