package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		m int
		n int
		x int
	}{
		{
			2,
			2,
			2,
		},
		{
			7,
			3,
			28,
		},
		{
			100,
			1,
			1,
		},
		{
			100,
			3,
			5050,
		},
	}

	for _, c := range cases {
		out := uniquePaths(c.m, c.n)
		if out != c.x {
			t.Fatalf("%d * %d\nexpect: %v\nactual: %v\n", c.m, c.n, c.x, out)
		}
	}
}
