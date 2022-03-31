package linkedlist

type Usecases interface {
	AddNode(v int)
	ToSlice(v int) []int
	FindHead() *LinkedList
}
