package main

import (
	"fmt"
	"time"
)

func main() {

	// canal <- {VALOR} = Envia dados para o canal
	// <- canal (esperando o canal receber o valor
	// mensagem := <- canal (Espera receber o valor do canal e armazena na variável mensagem

	fmt.Println("Go Routines Canais")

	canal := make(chan string)

	go escrever("Olá mundo", canal)

	//for{
	//	mensagem, aberto := <-canal
	//	fmt.Println(mensagem)
	//
	//	if !aberto{
	//		break
	//	}
	//}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
