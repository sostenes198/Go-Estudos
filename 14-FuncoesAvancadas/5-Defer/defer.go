package main

import "fmt"

func main() {
	fmt.Println("Defer")
	result := alunoEstaAprovado(5, 5)
	fmt.Println(result)
}

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	defer fmt.Println("Média calculada. Resultado será retornado!")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}
