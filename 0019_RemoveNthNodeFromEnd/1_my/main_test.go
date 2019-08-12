package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		input  []int
		n      int
		expect []int
	}{
		{
			[]int{1, 2},
			2,
			[]int{2},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[]int{1, 2, 3, 4},
		},
	}

	for _, c := range cases {
		// constrcut list
		head := &ListNode{Val: c.input[0]}
		node := head
		for _, i := range c.input[1:] {
			newNode := &ListNode{Val: i}
			node.Next = newNode
			node = newNode
		}

		out := removeNthFromEnd(head, c.n)

		// convert output list to slice
		slice := []int{}
		for out != nil {
			slice = append(slice, out.Val)
			out = out.Next
		}

		if len(slice) != len(c.expect) {
			t.Fatalf("\nexpect: %v\noutput: %v\n", c.expect, slice)
		}

		for i := range slice {
			if slice[i] != c.expect[i] {
				t.Fatalf("\nexpect: %v\noutput: %v\n", c.expect, slice)
			}
		}
	}
}
