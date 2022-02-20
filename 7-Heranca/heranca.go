// Só que não :)
package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	estudante1 := estudante{pessoa{"Soso", 26}, "Eng", "UEMG"}
	fmt.Println(estudante1)
}
