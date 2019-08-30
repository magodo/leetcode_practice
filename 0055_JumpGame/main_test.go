package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		input []int
		ok    bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
		{
			[]int{1, 2, 3},
			true,
		},
	}

	for _, c := range cases {
		out := canJump(c.input)
		if out != c.ok {
			t.Fatalf("%v\nexpect: %v\nactual: %v\n", c.input, c.ok, out)
		}
	}
}
