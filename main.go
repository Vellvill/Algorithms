package main

import (
	"algorithms/linkedlist"
	"fmt"
)

func main() {
	ll := linkedlist.CreateList(1)
	for i := 0; i <= 10; i++ {
		ll.AddNode(i)
	}
	u := ll.Reverse()
	fmt.Println(u.ToSlice())
}
