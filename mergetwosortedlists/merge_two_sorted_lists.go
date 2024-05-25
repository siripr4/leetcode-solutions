package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// mergeTwoLists()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var mergedList *ListNode
	var mergedListTail *ListNode

	for list1 != nil && list2 != nil {
		currentNode := &ListNode{}
		if list1.Val <= list2.Val {
			currentNode.Val = list1.Val
			list1 = list1.Next
		} else {
			currentNode.Val = list2.Val
			list2 = list2.Next
		}
		if mergedList == nil {
			mergedList = currentNode
			mergedListTail = mergedList
		}
		mergedListTail.Next = currentNode
		mergedListTail = mergedListTail.Next
	}

	if list1 != nil {
		mergedListTail.Next = list1
	}
	if list2 != nil {
		mergedListTail.Next = list2
	}
	return mergedList
}
