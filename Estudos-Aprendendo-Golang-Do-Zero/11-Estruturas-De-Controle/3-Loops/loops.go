package main

import (
	"fmt"
)

func main() {
	//i := 0
	//for i < 10{
	//	time.Sleep(time.Second)
	//	fmt.Println("Incrementando I", i)
	//	i++
	//}

	//for j := 0; j <10; j++{
	//	time.Sleep(time.Second)
	//	fmt.Println("Incrementando J", j)
	//}

	nomes := []string{"Joao", "Soso", "Joaquim"}

	for indice, nome := range nomes{
		fmt.Println(indice)
		fmt.Println(nome)
	}
	for indice, letra := range "Palavra" {
		fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "Ed",
		"idade": "28",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}
}
