package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		input string
		ok    bool
	}{
		{
			"((",
			false,
		},
		{
			"[]{}",
			true,
		},
		{
			"[{}]",
			true,
		},
		{
			"[{]",
			false,
		},
	}

	for _, c := range cases {
		out := isValid(c.input)
		if out != c.ok {
			t.Fatalf("%v expect to be %v, but got %v", c.input, c.ok, out)
		}
	}
}
