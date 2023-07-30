package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{6, 7, 8, 9, 10}
	fmt.Println(array3)

	slice1 := []int{11, 12, 13, 14, 15}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1), reflect.TypeOf(array3))

	slice1 = append(slice1, 16, 17, 18)
	fmt.Println(slice1)

	slice2 := array2[0:3]
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	var ponteiro = &slice1[0]
	var valor = *ponteiro
	println(&slice3[0])
	println(*ponteiro)
	println(ponteiro)
	println(valor)
}
