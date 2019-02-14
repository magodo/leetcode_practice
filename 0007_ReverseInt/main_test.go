package main

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		x int
		y int
	}{
		{
			120,
			21,
		},
		{
			123,
			321,
		},
		{
			-120,
			-21,
		},
		{
			-123,
			-321,
		},
		{
			1534236469,
			0,
		},
		{
			2147483648,
			0,
		},
		{
			-2147483648,
			0,
		},
	}

	for _, c := range cases {
		fmt.Println("---")
		out := reverse(c.x)
		if out != c.y {
			t.Fatalf("\nx: %d\ny: %d (expected: %d)\n", c.x, out, c.y)
		}
	}
}
