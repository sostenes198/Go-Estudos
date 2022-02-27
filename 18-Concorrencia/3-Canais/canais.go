package main

import (
	"fmt"
	"time"
)

func main() {
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
