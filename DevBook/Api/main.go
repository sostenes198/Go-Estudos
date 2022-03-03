package main

import (
	"devbook/src/config"
	"devbook/src/database"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// load .env file
	config.Load()

	// Database
	if err := database.Open(); err != nil{
		log.Fatal(err)
	}
	if err := database.Open(); err != nil{
		log.Fatal(err)
	}
	if err := database.Open(); err != nil{
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Printf("Rodando API na porta %d\n", config.Port)
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
