package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	if head.Next == nil {
		return nil
	}

	lnode := head
	rnode := head.Next.Next
	tnode := head
	for i := 0; i < n; i++ {
		tnode = tnode.Next
	}

	// n equals to length of list, just return head.Next
	if tnode == nil {
		return head.Next
	}

	for tnode.Next != nil {
		lnode, tnode = lnode.Next, tnode.Next
		if rnode != nil {
			rnode = rnode.Next
		}
	}
	lnode.Next = rnode
	return head
}
