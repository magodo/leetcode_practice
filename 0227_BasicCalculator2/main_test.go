package foo

import "testing"

func TestCalculator(t *testing.T) {
	cases := []struct {
		in     string
		expect int
	}{
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
	}

	for i, tt := range cases {
		got := calculate(tt.in)
		if got != tt.expect {
			t.Fatalf("%02d expected=%d, got=%d", i, tt.expect, got)
		}
	}

}
