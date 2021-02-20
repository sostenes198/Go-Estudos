package main

import "fmt"

func somar(n1 int32, n2 int32) int32{
	return n1+n2
}

func calculosMatematicos(n1, n2 int32) (int32, int32) {
	soma := n1+n2
	substracao := n1-n2
	return soma, substracao
}

func main(){
	fmt.Println(somar(1, 2))

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	f("Texto da função")

	resultadosCalculosSoma, _  := calculosMatematicos(10, 5)
	fmt.Println(resultadosCalculosSoma)
}
