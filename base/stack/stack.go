package stack

import (
    "algorithms/base/binary_tree"
)

type T interface {
    *binary_tree.BinaryTree | *binary_tree.BinaryNode
}

//LIFO
type stack[Allowed T] struct {
    s []Allowed
}

func NewQueue[Allowed T](el Allowed) *stack[Allowed] {
    q := new(stack[Allowed])
    q.s = append(q.s, el)
    return q
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
