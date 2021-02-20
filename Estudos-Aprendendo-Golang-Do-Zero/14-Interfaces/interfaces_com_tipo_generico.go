package main

import "fmt"

func generica(interF interface{}){
	fmt.Println(interF)
}

func main() {
	fmt.Println("Interfaces com tipo genérico")
	generica(1)
	generica("Awe")
	generica(true)
}
