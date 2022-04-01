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
	ll.Reverse()
	fmt.Println(ll)
}
