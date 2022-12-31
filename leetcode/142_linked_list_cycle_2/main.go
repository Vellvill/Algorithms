package main

import (
	"algorithms/base/linkedlist"
	"algorithms/base/random"
	"fmt"
)

func main() {
	fmt.Println(detectCycle(random.GetList(10, 5).MakeCycle()).Val)
	fmt.Println(detectCycle(random.GetList(10, 5)))
}

func detectCycle(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil {
		return nil
	}
	for slow, fast := head, head; ; {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			slow = head
			for slow != head {
				slow = slow.Next
			}
			return slow
		}
	}
}
