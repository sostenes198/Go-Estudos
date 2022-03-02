package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	soso := usuario{"Sóstenes", 26, endereco{"rua dos bobos", 0}}
	fmt.Println(soso)

	raquel := usuario{nome: "Raquel"}
	fmt.Println(raquel)
}
