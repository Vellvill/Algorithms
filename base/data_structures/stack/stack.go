package stack

import (
	"algorithms/base/data_structures/binary_tree"
	"algorithms/base/data_structures/custom_types"
)

type T interface {
	*binary_tree.BinaryTree | *binary_tree.BinaryNode | *custom_types.Pair
}

//LIFO
type stack[Allowed T] struct {
	s []Allowed
}

func NewStack[Allowed T](el Allowed) *stack[Allowed] {
	q := new(stack[Allowed])
	q.s = append(q.s, el)
	return q
}

func NewClearStack[Allowed T]() *stack[Allowed] {
	return new(stack[Allowed])
}

func (q *stack[Allowed]) Push(a Allowed) {
	q.s = append(q.s, a)
}

func (q *stack[Allowed]) Pop() (r Allowed) {
	r, q.s = q.s[len(q.s)-1], q.s[:len(q.s)-1]
	return r
}

func (q *stack[Allowed]) GetLen() int {
	return len(q.s)
}
