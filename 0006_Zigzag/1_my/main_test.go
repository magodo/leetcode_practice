package main

import "testing"

func TestAll(t *testing.T) {
	cases := []struct {
		in  string
		n   int
		out string
	}{
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		{
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		{
			"a",
			2,
			"a",
		},
		{
			"a",
			1,
			"a",
		},
	}

	for _, c := range cases {
		out := convert(c.in, c.n)
		if out != c.out {
			t.Fatalf("\nin: %s\nn: %d\nout: %s(expected: %s)\n", c.in, c.n, out, c.out)
		}
	}
}
