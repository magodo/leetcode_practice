package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToSlice() []int {
	out := []int{}
	for l != nil {
		out = append(out, l.Val)
		l = l.Next
	}
	return out
}

func FromSlice(slice []int) *ListNode {
	nodeBeforeHead := &ListNode{}
	l := nodeBeforeHead
	for _, i := range slice {
		l.Next = &ListNode{Val: i}
		l = l.Next
	}
	return nodeBeforeHead.Next
}

func mergeOne(l1 *ListNode, l2 *ListNode, lm *ListNode) (newL1, newL2 *ListNode, newLm *ListNode) {
	if l1.Val < l2.Val {
		newNode := &ListNode{Val: l1.Val}
		lm.Next = newNode
		return l1.Next, l2, lm.Next
	}
	newNode := &ListNode{Val: l2.Val}
	lm.Next = newNode
	return l1, l2.Next, lm.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	nodeBeforeHead := &ListNode{}
	l := nodeBeforeHead
	for l1 != nil && l2 != nil {
		l1, l2, l = mergeOne(l1, l2, l)
	}

	var ltail *ListNode
	if l1 == nil {
		ltail = l2
	} else {
		ltail = l1
	}
	for ltail != nil {
		l.Next = &ListNode{Val: ltail.Val}
		ltail, l = ltail.Next, l.Next
	}

	return nodeBeforeHead.Next
}
