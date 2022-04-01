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
			n.next = new
			break
		}
	}
}

func (l *LinkedList) ToSlice() []int {
	out := make([]int, 0)
	for n := l; n != nil; n = n.next {
		out = append(out, n.value)
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

func (l *LinkedList) Delete(v int) {
	for n := l; n != nil; n = n.next {
		if n.value == v {
			n.value = n.next.value
			n.next = n.next.next
		}
	}
}

func (l *LinkedList) Reverse() *LinkedList {

	cur := l
	var prev, next *LinkedList

	for cur != nil {
		//копируем указатель на сл элемент
		next = cur.next

		//разворачиваем следующий указатель
		cur.next = prev

		//идём дальше по листу
		prev = cur
		cur = next
	}

	return prev

}

func (l *LinkedList) AddNodeByValue(WhichToAdd, Adder int) {
	for n := l; n != nil; n = n.next {
		if n.value == WhichToAdd {
			new := CreateList(Adder)
			new.next = n.next
			n.next = new
			break
		}
	}
}

func (l *LinkedList) AddNodeToBegin(v int) {
	new := CreateList(v)
	new.next = l
}

