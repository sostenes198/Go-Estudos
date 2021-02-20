package main

import (
	"fmt"
	"time"
)

func main(){
	// CONCORRENCIA != PARALELISMO
	fmt.Println("goroutines")
	go escrever("Olá mundo")
	escrever("Programando em GO")
}

func escrever(texto string){
	for{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}