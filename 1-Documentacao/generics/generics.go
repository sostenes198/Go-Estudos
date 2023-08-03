package main

import "fmt"

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	a := A{}
	runnerA := Test[A]{Runner: a}
	runnerA.Runner.Run()

	b := B{}
	runnerB := Test[B]{Runner: b}
	runnerB.Runner.Run()
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type IInterface interface {
	Run()
}

type A struct {
}

func (a A) Run() {
	fmt.Println("A")
}

type B struct {
}

func (b B) Run() {
	fmt.Println("B")
}

type Test[V IInterface] struct {
	Runner V
}

type Number interface {
	int64 | float64

	Runner(as Number)
}
