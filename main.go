package main

import (
	"algorithms/array"
	"fmt"
)

func main() {
	m := array.MakeArray()
	m.Append(4, 6, 7, 3, 4, 5, 6)
	m.Reverse()
	fmt.Println(m)
	m.RemoveDuplicates()
	fmt.Println(m)
}
