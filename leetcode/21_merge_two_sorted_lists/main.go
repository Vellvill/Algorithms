package main

//enable for random structs
import (
	"algorithms/base/linkedlist"
	. "algorithms/base/random"
)

//21_merge_two_sorted_lists
func main() {
	l1 := GetSortedList(GetIntInRange(1, 5), 10, 100)
	l2 := GetSortedList(GetIntInRange(1, 5), 10, 100)
	l1.Print()
	l2.Print()
	mergeTwoLists(l1, l2).Print()
}

func mergeTwoLists(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	l1, l2 := list1, list2

	dummy := &linkedlist.ListNode{}
	n := dummy

	for ; l1 != nil && l2 != nil; n = n.Next {
		if l1.Val < l2.Val {
			n.Next = l1
			l1 = l1.Next
		} else {
			n.Next = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		n.Next = l1
	} else if l2 != nil {
		n.Next = l2
	}

	return dummy.Next
}
