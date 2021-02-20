package main

import (
	"15-Aplicacao-Linha-De-Comando/app"
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if error := aplicacao.Run(os.Args); error != nil{
		log.Fatal(error)
	}
}
