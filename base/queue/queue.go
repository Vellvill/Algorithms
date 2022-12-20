package queue

import (
    "algorithms/base/binary_tree"
)

type T interface {
    *binary_tree.BinaryTree | *binary_tree.BinaryNode
}

type queue[Allowed T] struct {
    q []Allowed
}

func NewQueue[Allowed T](el Allowed) *queue[Allowed] {
    q := new(queue[Allowed])
    q.q = append(q.q, el)
    return q
}

func (q *queue[Allowed]) Push(a Allowed) {
    q.q = append(q.q, a)
}

func (q *queue[Allowed]) Pop() Allowed {
    r := q.q[0]
    q.q = q.q[1:]
    return r
}

func (q *queue[Allowed]) GetLen() int {
    return len(q.q)
}
