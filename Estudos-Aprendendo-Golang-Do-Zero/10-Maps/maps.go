package main

import "fmt"

func main(){
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome" : "Soso",
		"sobrenome": "Souza",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome" : {
			"primeiro": "Soso",
			"segundo": "Souza",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Genero",
	}
	fmt.Println(usuario2)
}
