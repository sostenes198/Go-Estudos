package Formas

import (
	"math"
)

type Forma interface{
	Area() float64
}

type Retangulo struct{
	Altura float64
	Largural float64
}

func(r Retangulo) Area() float64{
	return r.Altura * r.Largural
}

type Circulo struct{
	Raio float64
}

func(c Circulo) Area() float64{
	return math.Pi * math.Pow(c.Raio, 2)
}
