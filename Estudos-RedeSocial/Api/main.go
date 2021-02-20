package main

import (
	"Api/src/config"
	"Api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()
	fmt.Println("Rodando API")
	log.Fatal(http.ListenAndServe(":5000", router.Gerar()))
}
