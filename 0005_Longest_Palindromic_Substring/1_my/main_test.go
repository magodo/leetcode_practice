package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		s   string
		sub string
	}{
		{
			"baaab",
			"baaab",
		},
		{
			"aabad",
			"aba",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"a",
			"a",
		},
		{
			"aa",
			"aa",
		},
	}

	for _, c := range cases {
		output := longestPalindrome(c.s)
		if output != c.sub {
			t.Fatalf("\nstring: %s\noutput: %s(expect: %s)\n", c.s, output, c.sub)
		}
	}
}
