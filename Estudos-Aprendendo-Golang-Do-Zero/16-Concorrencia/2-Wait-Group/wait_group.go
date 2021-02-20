package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Group")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escreverWaitGroup("Olá mundo")
		waitGroup.Done()
	}()

	go func() {
		escreverWaitGroup("Go Routinesss !!! CONCORRENCIA")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escreverWaitGroup(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
