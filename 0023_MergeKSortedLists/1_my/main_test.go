package main

import (
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		inputs [][]int
		output []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{},
			},
			[]int{1, 2, 3},
		},
		{
			[][]int{
				[]int{},
				[]int{},
			},
			[]int{},
		},
		{
			[][]int{
				[]int{1, 4, 5},
				[]int{1, 3, 4},
				[]int{2, 6},
			},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}

	for _, c := range cases {
		// c't input
		listNodes := []*ListNode{}
		for _, input := range c.inputs {
			listNodes = append(listNodes, FromSlice(input))
		}

		out := mergeKLists(listNodes).ToSlice()
		if !reflect.DeepEqual(out, c.output) {
			t.Fatalf("\nexpect: %v\noutput: %v\n", c.output, out)
		}
	}
}
