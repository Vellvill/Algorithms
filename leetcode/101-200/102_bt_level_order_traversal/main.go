package main

import (
	"algorithms/base/core/random"
	"algorithms/base/data_structures/binary_tree"
	"algorithms/base/data_structures/queue"
	"fmt"
)

func main() {
	fmt.Println(levelOrder(random.GetIntTree(10, 10).Root))
}

func levelOrder(root *binary_tree.BinaryNode) [][]int {
	if root == nil {
		return nil
	}
	qu, res := queue.NewQueue(root), make([][]int, 0)
	for qu.GetLen() > 0 {
		qs := qu.GetLen()
		var level = make([]int, 0)
		for i := 0; i < qs; i++ {
			node := qu.Pop()

			level = append(level, node.Data)

			if node.LeftNode != nil {
				qu.Push(node.LeftNode)
			}
			if node.RightNode != nil {
				qu.Push(node.RightNode)
			}
		}
		res = append(res, level)
	}
	return res
}
