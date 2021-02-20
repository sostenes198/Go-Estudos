package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Canais")

	canal := make(chan string)
	go escreverCanais("Olá mundo", canal)

	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	//
	//}

	for mensagem := range canal{
		fmt.Println(mensagem)
	}
}

func escreverCanais(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
