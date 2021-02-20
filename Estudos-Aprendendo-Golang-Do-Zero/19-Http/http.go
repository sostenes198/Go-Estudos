package main

import (
	"log"
	"net/http"
)

func main(){
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ UMA REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(responseWriter http.ResponseWriter, request *http.Request){
		responseWriter.Write([]byte("Olá mundo"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
