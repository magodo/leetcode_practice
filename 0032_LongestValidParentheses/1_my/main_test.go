package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		input string
		n     int
	}{
		{
			")()()",
			4,
		},
		{
			"(())",
			4,
		},
		{
			")(())",
			4,
		},
		{
			")()())",
			4,
		},
		{
			")()())()()(",
			4,
		},
	}

	for _, c := range cases {
		out := longestValidParentheses(c.input)
		if out != c.n {
			t.Fatalf("\nexpect: %v\ngot: %v\n", c.n, out)
		}
	}
}
