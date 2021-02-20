package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main(){
	fmt.Println("Escrevendo arquivo main")
	auxiliar.Escrever()
	error := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(error)
}
