package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var nodePointers []*ListNode

	for head.Next != nil {
		for _, nodePointer := range nodePointers {
			if head == nodePointer {
				return true
			}
		}
		nodePointers = append(nodePointers, head)
		head = head.Next
	}

	return false
}

func hasCycleTwoPointers(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
