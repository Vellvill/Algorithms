package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(v int) *ListNode {
	return &ListNode{Val: v, Next: nil}
}

func (l *ListNode) AddNode(v int) {
	for n := l; n != nil; n = n.Next {
		if n.Next == nil {
			new := CreateList(v)
			n.Next = new
			break
		}
	}
}

func (l *ListNode) MakeCycle() *ListNode {
	for n := l; n != nil; n = n.Next {
		if n.Next == nil {
			n.Next = l.FindHead()
			break
		}
	}
	return l
}

func (l *ListNode) ToSlice() []int {
	out := make([]int, 0)
	for n := l; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func (l *ListNode) FindHead() *ListNode {
	var Head *ListNode
	for n := l; n != nil; n = n.Next {
		Head = n
	}
	return Head
}

func (l *ListNode) Delete(v int) {
	for n := l; n != nil; n = n.Next {
		if n.Val == v {
			n.Val = n.Next.Val
			n.Next = n.Next.Next
		}
	}
}

func (l *ListNode) Reverse() *ListNode {

	cur := l
	var prev, next *ListNode

	for cur != nil {
		//копируем указатель на сл элемент
		next = cur.Next

		//разворачиваем следующий указатель
		cur.Next = prev

		//идём дальше по листу
		prev = cur
		cur = next
	}

	return prev

}

func (l *ListNode) AddNodeByValue(WhichToAdd, Adder int) {
	for n := l; n != nil; n = n.Next {
		if n.Val == WhichToAdd {
			new := CreateList(Adder)
			new.Next = n.Next
			n.Next = new
			break
		}
	}
}

func (l *ListNode) AddNodeToBegin(v int) {
	new := CreateList(v)
	new.Next = l
}

func (l *ListNode) Print() {
	fmt.Print("\n| ")
	for n := l; n != nil; n = n.Next {
		fmt.Print(n.Val, " ")
		if n.Next != nil {
			fmt.Print("--> ")
		} else {
			fmt.Print("|")
		}
	}
}
