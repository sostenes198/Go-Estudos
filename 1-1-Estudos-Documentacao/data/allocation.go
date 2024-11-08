package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	//p := new(SyncedBuffer) // type *SyncedBuffer
	//var v SyncedBuffer     // type  SyncedBuffer
	//
	//x := p.lock
	//y := p.buffer
	//
	//fmt.Println(x)
	//fmt.Println(y)
	//
	//x = v.lock
	//y = v.buffer
	//
	//fmt.Println(x)
	//fmt.Println(y)
	//
	//fmt.Println(p)
	//fmt.Println(v)

	//var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
	//var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints
	//
	//fmt.Println(p)
	//fmt.Println(v)
	//
	//// Unnecessarily complex:
	//*p = make([]int, 100, 100)
	//
	//fmt.Println(p)

	//array := [...]float64{7.0, 8.5, 9.1}
	//x := Sum(&array) // Note the explicit address-of operator
	//
	//fmt.Println(x)

	//integers := []int{10, 20, 30}
	//fmt.Println(&integers)
	//fmt.Println(&integers[0])
	//ReceiveArray(integers)
	//fmt.Println(&integers)

	x := make([]int, 10)
	fmt.Println(x)
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func ReceiveArray(integers []int) {
	integers = append(integers, 55)
	fmt.Println(&integers[0])
	fmt.Println(&integers)
}
