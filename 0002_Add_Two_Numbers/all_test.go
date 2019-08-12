package main

import (
	"fmt"
	"testing"
)

func genListNode(nums ...int) *ListNode {
	var root, prev *ListNode
	for _, i := range nums {
		if root == nil {
			root = &ListNode{i, nil}
			prev = root
		} else {
			node := &ListNode{i, nil}
			prev.Next = node
			prev = node
		}
	}
	return root
}

func compareList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	if l1 == nil && l2 == nil {
		return true
	}
	return false
}

func dumpList(l *ListNode) {
	if l == nil {
		return
	}
	fmt.Printf("%d", l.Val)
	l = l.Next
	for l != nil {
		fmt.Printf(" -> %d", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func TestAddTwoNumbers(t *testing.T) {
	tables := []struct {
		l1 *ListNode
		l2 *ListNode
		lo *ListNode
	}{
		{
			l1: genListNode(0, 0, 1),
			l2: genListNode(0, 1),
			lo: genListNode(0, 1, 1),
		},
		{
			l1: genListNode(0),
			l2: genListNode(0, 1, 1),
			lo: genListNode(0, 1, 1),
		},
		{
			l1: genListNode(0),
			l2: genListNode(0),
			lo: genListNode(0),
		},
		{
			l1: genListNode(9, 9, 1),
			l2: genListNode(1),
			lo: genListNode(0, 0, 2),
		},
		{
			l1: genListNode(9),
			l2: genListNode(1),
			lo: genListNode(0, 1),
		},
	}

	for _, table := range tables {
		lo := addTwoNumbersV2(table.l1, table.l2)
		if !compareList(lo, table.lo) {
			dumpList(table.l1)
			dumpList(table.l2)
			dumpList(lo)
			dumpList(table.lo)
			t.Fatal("failure")
		}
	}
}
