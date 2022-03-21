package main

import (
	"fmt"
	"sort"
)

func main() {

}

type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
