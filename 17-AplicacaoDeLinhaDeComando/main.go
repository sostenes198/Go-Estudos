package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if error := aplicacao.Run(os.Args); error != nil {
		log.Fatalln(error)
	}

}
