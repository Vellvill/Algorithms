package linkedlist

type LinkedList struct {
	value int
	next  *LinkedList
}

func CreateList(v int) *LinkedList {
	return &LinkedList{value: v, next: nil}
}

func (l *LinkedList) AddNode(v int) {
	for n := l; n != nil; n = n.next {
		if n.next == nil {
			new := CreateList(v)
			l.next = new
			break
		}
	}
}

func (l *LinkedList) ToSlice(v int) []int {
	out := make([]int, 0)
	for n := l; n != nil; n = n.next {
		out = append(out, v)
	}
	return out
}

func (l *LinkedList) FindHead() *LinkedList {
	var Head *LinkedList
	for n := l; n != nil; n = n.next {
		Head = n
	}
	return Head
}
