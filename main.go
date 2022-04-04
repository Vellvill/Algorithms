package main

import (
	"algorithms/binary_tree"
	"os"
)

func main() {
	tree := binary_tree.NewTree()
	tree.Insert(100)
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(75)
	tree.Insert(175)
	binary_tree.PrintTree(os.Stdout, tree.Root, 0, 'M')
}
