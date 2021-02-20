package main

import (
	"fmt"
	"math"
)

type forma interface{
	area() float64
}

func escreverArea(f forma){
	fmt.Println("A área da forma é ", f.area())
}

type retangulo struct{
	altura float64
	largural float64
}

func(r retangulo) area() float64{
	return r.altura * r.largural
}

type circulo struct{
	raio float64
}

func(c circulo) area() float64{
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	fmt.Println("Interfaces")

	r := retangulo{10, 15}
	c := circulo{100}
	escreverArea(r)
	escreverArea(c)
}
