package linkedlist

type Usecases interface {
	AddNode(v int)
	ToSlice() []int
	FindHead() *LinkedList
	Delete(v int)
}
