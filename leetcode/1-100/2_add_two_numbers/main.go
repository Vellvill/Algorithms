package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"algorithms/base/data_structures/linkedlist"
)

//2_add_two_numbers
func main() {
	l1 := linkedlist.CreateList(9)
	l1.AddNode(9)
	l1.AddNode(9)
	l2 := linkedlist.CreateList(9)
	l2.AddNode(9)
	l2.AddNode(9)
	addTwoNumbers(l1, l2).Print()
}

func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	r1, r2 := reverse(l1), reverse(l2)

	var dummy = &linkedlist.ListNode{}
	var overHead int

	for n1, n2, d := r1, r2, dummy; n1 != nil && n2 != nil; n1, n2, d = n1.Next, n2.Next, d.Next {
		sum := n1.Val + n2.Val + overHead
		if overHead > 0 {
			overHead = 0
		}
		if sum >= 10 {
			sum = sum % 10
			overHead += 1
		}
		d.Next = &linkedlist.ListNode{
			Val: sum,
		}
	}

	return reverse(dummy.Next)
}

func reverse(l *linkedlist.ListNode) *linkedlist.ListNode {
	var prev, next *linkedlist.ListNode
	var curr = l
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
