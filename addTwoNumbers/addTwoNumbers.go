package addTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := &ListNode{}
	resultTail := result
	for {
		if l1 != nil && l2 != nil {
			if l1.Val + l2.Val + carry >= 10 {
				resultTail.Val = (l1.Val + l2.Val+carry)%10
				carry = 1
			} else {
				resultTail.Val = l1.Val + l2.Val + carry
				carry = 0
			}
			if l1.Next != nil && l2.Next == nil {
				l2.Next = &ListNode{Val:0, Next:nil}
			}
			if l2.Next != nil && l1.Next == nil {
				l1.Next = &ListNode{Val:0, Next:nil}
			}
			if l1.Next != nil || l2.Next != nil {
				resultTail.Next = &ListNode{}
				resultTail = resultTail.Next
			} else {
				if carry == 1 {
					resultTail.Next = &ListNode{Val:1}
					resultTail = resultTail.Next
				}
			}
			l2 = l2.Next
			l1 = l1.Next
		} else {
			break
		}

	}
	return result
}