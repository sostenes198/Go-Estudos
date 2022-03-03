package main

import "fmt"

func main() {
	fmt.Println("Panic e Recover")

	alunoEstaAprovado(6, 6)
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXTAMENTE 6!")
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada com sucesso.")
	}
}
