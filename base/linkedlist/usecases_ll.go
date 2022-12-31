package linkedlist

type Usecases interface {
	AddNode(v int)
	ToSlice() []int
	FindHead() *ListNode
	Delete(v int)
	Reverse() *ListNode
	AddNodeByValue(WhichToAdd, Adder int)
	AddNodeToBegin(v int)
}
