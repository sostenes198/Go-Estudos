package main

import "fmt"

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
}

type ByteSlice []byte

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// Body as above, without the return.
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	*p = slice
	return len(data), nil
}
