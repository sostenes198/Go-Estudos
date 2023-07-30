package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando usuário %s em memória", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func (u usuario) fazerAniversario2() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")

	usuario1 := usuario{"Usuário 1", 20}
	usuario2 := usuario{"Usuário 2", 20}
	usuario1.salvar()
	usuario2.salvar()

	usuario3 := usuario{"Usuário 3", 20}
	usuario3.fazerAniversario()
	fmt.Println(usuario3.idade)

	usuario4 := usuario{"Usuário 4", 15}
	usuario4.fazerAniversario2()
	fmt.Println(usuario4.idade)
}
