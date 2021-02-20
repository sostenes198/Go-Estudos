package main

import "fmt"

type Usuario struct {
	nome string
	idade int
	Endereco Endereco
}

type Endereco struct {
	logradouro string
	numero int
}

func main(){
	var u Usuario
	u.nome = "Sóstenes"
	u.idade = 25
	fmt.Println(u)

	endereco := Endereco{"Rua dos bobos", 0}

	usuario2 := Usuario{"Davi", 25, endereco}
	fmt.Println(usuario2)

	usuario3 := Usuario{idade: 25}
	fmt.Println(usuario3)
}
