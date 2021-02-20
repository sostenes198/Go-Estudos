package main

import "fmt"

type usuario struct{
	nome string
	idade int
}

func(u usuario) salvar(){
	fmt.Println("Salvandos dados do usuário", u.nome, " No banco de dados")
}

func(u usuario) maiorDeIdade() bool{
	return u.idade >= 18
}

func(u *usuario) fazerAniversario(){
	u.idade++
}

func main() {
	usuario1 := usuario{"Soso", 25}
	usuario1.salvar()

	usuario2 := usuario{"Jose", 25}
	usuario2.salvar()

	fmt.Println(usuario1.maiorDeIdade())

	fmt.Println(usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
