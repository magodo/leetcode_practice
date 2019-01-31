package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carray := 0
	var prev, root *ListNode

	for {
		// calculate current node's value, together with next carray
		val := carray
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		carray = val / 10
		val = val % 10

		node := &ListNode{
			Val:  val,
			Next: nil,
		}

		// update previous node's Next pointer
		if prev == nil {
			root = node
		} else {
			prev.Next = node
		}
		prev = node

		// iterate two lists
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			// consider the carray
			if carray != 0 {
				prev.Next = &ListNode{
					Val:  carray,
					Next: nil,
				}
			}
			return root
		}
	}
}
