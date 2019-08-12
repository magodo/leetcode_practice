package main

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode
	prev := &ListNode{}
	carry := 0

	for !(l1 == nil && l2 == nil && carry == 0) {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		val := sum % 10
		carry = sum / 10

		prev.Next = &ListNode{
			Val: val,
		}
		prev = prev.Next
		if root == nil {
			root = prev
		}
	}
	prev.Next = nil
	return root
}
