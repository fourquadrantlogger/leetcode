package add_two_numbers

import "leetcode/common"

func AddTwoNumbers(l1, l2 *common.ListNode) (head *common.ListNode) {
	var tail *common.ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &common.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &common.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &common.ListNode{Val: carry}
	}
	return
}
