package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		nums   []int
		length int
	}{
		{
			[]int{1, 2, 3, 3, 2, 5, 6, 5, 7, 8},
			6,
		},
		{
			[]int{2, 6, 4, 8, 10, 9, 15},
			5,
		},
		{
			[]int{5, 4, 4, 3, 2, 1},
			6,
		},
		{
			[]int{1, 2, 2, 3},
			0,
		},
		{
			[]int{2, 2},
			0,
		},
		{
			[]int{1, 3, 2, 2, 2},
			4,
		},
		{
			[]int{1, 3, 2, 4, 5},
			2,
		},
	}

	for _, c := range cases {
		println(findUnsortedSubarray(c.nums))
		/*
			out := findUnsortedSubarray(c.nums)
			if out != c.length {
				t.Fatalf("expect: %d, actual: %d", c.length, out)
			}
		*/
	}
}
