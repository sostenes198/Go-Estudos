package main

import (
	"devbook/src/config"
	"devbook/src/database"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

//func init() {
//	key := make([]byte, 64)
//
//	if _, err := rand.Read(key); err != nil {
//		log.Fatal(err)
//	}
//
//	keyBase64 := base64.StdEncoding.EncodeToString(key)
//
//	fmt.Println(keyBase64)
//}

func main() {
	// load .env file
	config.Load()

	// Database
	if err := database.Open(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	fmt.Printf("Rodando API na porta %d\n", config.Port)
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
