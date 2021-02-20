package main

import "fmt"

func funcao1(){
	fmt.Println("Executando função 1")
}

func funcao2(){
	fmt.Println("Executando função 2")
}

func alunoAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se aluno esta aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {return true}
	return false
}

func main(){
	fmt.Println("Clausula defer")

	//// DEFER = adiar
	//defer funcao1()
	//funcao2()

	fmt.Println(alunoAprovado(7, 8))
}
