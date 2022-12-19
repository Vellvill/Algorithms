package linkedlist

type Usecases interface {
	AddNode(v int)
	ToSlice() []int
	FindHead() *LinkedList
	Delete(v int)
	Reverse() *LinkedList
	AddNodeByValue(WhichToAdd, Adder int)
	AddNodeToBegin(v int)
}
