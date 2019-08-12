package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func (l *ListNode) ToSlice() []int {
//	out := []int{}
//	for l != nil {
//		out = append(out, l.Val)
//		l = l.Next
//	}
//	return out
//}
//
//func FromSlice(slice []int) *ListNode {
//	nodeBeforeHead := &ListNode{}
//	l := nodeBeforeHead
//	for _, i := range slice {
//		l.Next = &ListNode{Val: i}
//		l = l.Next
//	}
//	return nodeBeforeHead.Next
//}

func mergeKLists(lists []*ListNode) *ListNode {
	nodeBeforeRoot := &ListNode{}
	node := nodeBeforeRoot
	var minNodeIdx int
	for {
		minNodeIdx = -1
		for idx := range lists {
			if lists[idx] != nil {
				if minNodeIdx == -1 || lists[idx].Val < lists[minNodeIdx].Val {
					minNodeIdx = idx
				}
			}
		}
		if minNodeIdx == -1 {
			break
		}
		node.Next = &ListNode{Val: lists[minNodeIdx].Val}
		node = node.Next
		lists[minNodeIdx] = lists[minNodeIdx].Next
	}
	return nodeBeforeRoot.Next
}
