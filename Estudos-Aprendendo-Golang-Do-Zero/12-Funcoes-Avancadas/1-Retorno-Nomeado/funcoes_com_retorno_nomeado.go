package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main()  {
	fmt.Println("Funcoes com Retornos nomeados")
	resultadoSoma, resultadoSubtracao := calculosMatematicos(1, 5)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
