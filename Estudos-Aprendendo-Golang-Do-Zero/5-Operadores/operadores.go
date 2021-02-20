package main

import "fmt"

func main(){
	// Aritimeticos
	soma := 1 +2
	substracao := 1 -2
	divisao := 10 / 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, substracao, divisao, restoDaDivisao)

	// Atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// unirários
	numero := 10
	numero++
	numero+=10
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
}
