package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
}

type estudante struct {
	pessoa
	estuda bool
}

func main(){
	fmt.Println("Herança")

	p1 := pessoa{"Sostenes", "Gonçalves", 20}
	fmt.Println(p1)

	estudante1 := estudante{p1, true}
	fmt.Println(estudante1)

	estudante2 := estudante{pessoa{"Soso", "Souza", 25}, true}
	fmt.Println(estudante2)
}
