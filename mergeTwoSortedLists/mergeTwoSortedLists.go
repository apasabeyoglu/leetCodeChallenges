package mergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultTail := result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			resultTail.Next = l2
			break
		}

		if l2 == nil {
			resultTail.Next = l1
			break
		}

		if l1.Val < l2.Val {
			resultTail.Next = l1
			l1 = l1.Next
		} else {
			resultTail.Next = l2
			l2 = l2.Next
		}
		resultTail = resultTail.Next
	}
	return result.Next
}
