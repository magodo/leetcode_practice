package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		nums []int
		n    int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{},
			0,
		},
	}

	for _, c := range cases {
		out := lengthOfLIS(c.nums)
		if out != c.n {
			t.Fatalf("%v\nexpect: %v\noutput: %v\n", c.nums, c.n, out)
		}
	}
}
