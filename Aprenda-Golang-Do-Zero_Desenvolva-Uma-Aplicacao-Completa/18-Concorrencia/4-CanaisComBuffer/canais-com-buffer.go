package main

import "fmt"

func main() {
	fmt.Println("Canais com Buffer")

	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "Programando em Go"

	mensagem := <-canal
	mensagem2 := <-canal
	
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
