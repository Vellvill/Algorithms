package main

import (
	"algorithms/base/linkedlist"
	"algorithms/base/random"
	"fmt"
)

func main() {
	fmt.Println(hasCycle(random.GetList(10, 5).MakeCycle()))
	fmt.Println(hasCycle(random.GetList(10, 5)))
}

func hasCycle(head *linkedlist.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	for slow, fast := head, head.Next; fast != nil && fast.Next != nil; slow, fast = slow.Next, fast.Next.Next {
		if slow == fast {
			return true
		}
	}
	return false
}
