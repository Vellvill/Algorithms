package main

import (
	"algorithms/linkedlist"
	"fmt"
)

func main() {
	l := linkedlist.CreateList(1)
	l.AddNode(2)
	l.AddNode(3)

	head := l.FindHead()
	fmt.Println(head)
}
