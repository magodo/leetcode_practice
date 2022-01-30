package foo

import "testing"

func TestCalculator(t *testing.T) {
	cases := []struct {
		in     string
		expect int
	}{
		{
			"(7)-(0)+(4)",
			11,
		},
		{
			"(1+(4+5+2)-3)+(6+8)",
			23,
		},
		{
			"-2+1",
			-1,
		},
		{
			"(7)-(0)+(4)",
			11,
		},
		{
			"1+(-2)",
			-1,
		},
		{
			"1+1",
			2,
		},
		{
			"1+2+3",
			6,
		},
		{
			"1-1",
			0,
		},
		{
			"1+2*3-4",
			3,
		},
		{
			"5 +4*3/2- 1",
			10,
		},
		{
			"5 +3/2*4- 1",
			8,
		},
		{
			"5 +4*3**2/2- 1",
			22,
		},
	}

	for i, tt := range cases {
		got := calculate(tt.in)
		if got != tt.expect {
			t.Fatalf("%02d expected=%d, got=%d", i, tt.expect, got)
		}
	}

}
