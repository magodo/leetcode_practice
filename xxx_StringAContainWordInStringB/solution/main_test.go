package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		a  string
		b  string
		ok bool
	}{
		{
			a:  "ABCDEF",
			b:  "BABA",
			ok: true,
		},
		{
			a:  "ABCDEF",
			b:  "BABAZ",
			ok: false,
		},
	}

	for _, c := range cases {
		out := IsSubstring(c.a, c.b)
		if out != c.ok {
			t.Fatalf("%s\n%s\nexpect: %v\nactual: %v\n", c.a, c.b, c.ok, out)
		}
	}
}
