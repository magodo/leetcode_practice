package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		x  int
		ok bool
	}{
		{
			121,
			true,
		},
		{
			-121,
			false,
		},
		{
			11,
			true,
		},
		{
			1,
			true,
		},
		{
			0,
			true,
		},
	}

	for _, c := range cases {
		out := isPalindrome(c.x)
		if out != c.ok {
			t.Fatalf("%d: %v (expected %v)", c.x, out, c.ok)
		}
	}
}
