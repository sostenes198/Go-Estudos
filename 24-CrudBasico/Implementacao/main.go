package main

import (
	"fmt"
	"implementacao/entitity"
)

func main() {

	x := &entitity.User{Name: "asadasd", Email: "asudhasduh"}
	c := &entitity.User{Name: "asadasd", Email: "asudhasduh"}

	var a int
	var b *int

	a = 10
	b = &a

	fmt.Println(&x)
	fmt.Println(&c)

	fmt.Println(a, &b)

}
