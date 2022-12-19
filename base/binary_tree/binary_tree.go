package binary_tree

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	LeftNode  *BinaryNode
	RightNode *BinaryNode
	Data      int
}

type BinaryTree struct {
	Root *BinaryNode
}

func NewTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func NewNode(data int) *BinaryNode {
	return &BinaryNode{
		LeftNode:  nil,
		RightNode: nil,
		Data:      data,
	}
}

func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{
			Data:      data,
			LeftNode:  nil,
			RightNode: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

func (n *BinaryNode) Insert(data int) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.LeftNode == nil {
			n.LeftNode = NewNode(data)
		} else {
			n.LeftNode.Insert(data)
		}
	} else {
		if n.RightNode == nil {
			n.RightNode = NewNode(data)
		} else {
			n.RightNode.Insert(data)
		}
	}
}

func PrintTree(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	PrintTree(w, node.LeftNode, ns+2, 'L')
	PrintTree(w, node.RightNode, ns+2, 'R')
}
