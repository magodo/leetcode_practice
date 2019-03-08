package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		hights []int
		area   int
	}{
		{
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
		{
			[]int{2, 1, 5, 6, 6, 6},
			20,
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 2147483647},
			2147483647,
		},
		{
			[]int{5, 4, 4, 6, 3, 2, 9, 5, 4, 8, 1, 0, 0, 4, 7, 2},
			20,
		},
		{
			[]int{3, 6, 5, 7, 4, 8, 1, 0},
			20,
		},
	}

	for _, c := range cases {
		out := largestRectangleArea(c.hights)
		if out != c.area {
			t.Fatalf("expect: %d, got: %d", c.area, out)
		}
	}
}
