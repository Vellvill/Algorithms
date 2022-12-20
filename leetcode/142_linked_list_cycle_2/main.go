package main

import (
    "algorithms/base/linkedlist"
    "fmt"
    "algorithms/base/random"
)

func main() {
    fmt.Println(detectCycle(random.GetRandomList(10, 5).MakeCycle()).Value)
    fmt.Println(detectCycle(random.GetRandomList(10, 5)))
}

func detectCycle(head *linkedlist.LinkedList) *linkedlist.LinkedList {
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
