package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[int]map[string]string{
		1: {
			"nome":      "Soso",
			"sobrenome": "Souza",
		},
		2: {
			"nome":      "raquel",
			"sobrenome": "gontijo",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, 1)
	fmt.Println(usuario2)

	usuario2[3] = map[string]string{
		"nome": "Aleatório",
	}
	fmt.Println(usuario2)
}
