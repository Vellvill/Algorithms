package _26_invert_binary_tree

import (
	"algorithms/base/core/random"
	"algorithms/base/data_structures/binary_tree"
)

func main() {
	invertTree(random.GetIntTree(random.GetInt(10), random.GetInt(10)).Root)
}

func invertTree(root *binary_tree.BinaryNode) *binary_tree.BinaryNode {
	if root == nil {
		return nil
	}

	l := invertTree(root.LeftNode)
	r := invertTree(root.RightNode)

	root.LeftNode = r
	root.RightNode = l

	return root
}
