package main

import (
	"algorithms/base/core/random"
	"algorithms/base/data_structures/binary_tree"
	"fmt"
)

var (
	start = 1 << 32
)

func main() {
	fmt.Println(isValidBST(random.GetIntTree(2, 10).Root))
}

func isValidBST(root *binary_tree.BinaryNode) bool {
	return isValidBSTRecursion(root, start, start)
}

func isValidBSTRecursion(root *binary_tree.BinaryNode, min, max int) bool {
	if root == nil {
		return true
	} else if min != start && root.Data <= min || max != start && root.Data >= max {
		return false
	}
	return isValidBSTRecursion(root.LeftNode, min, root.Data) && isValidBSTRecursion(root.RightNode, root.Data, max)
}
