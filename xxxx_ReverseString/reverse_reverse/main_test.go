package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		s   []rune
		idx int
		out string
	}{
		{
			s:   []rune("12345"),
			idx: 2,
			out: "34512",
		},
	}

	for _, c := range cases {
		input := string(c.s)
		Reverse(c.s, c.idx)
		if string(c.s) != c.out {
			t.Fatalf("%s\nexpect: %s\nactual: %s\n", input, c.out, string(c.s))
		}
	}
}
