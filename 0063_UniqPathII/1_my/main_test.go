package main

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		grid [][]int
		n    int
	}{
		{
			[][]int{
				[]int{0, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			2,
		},
		{
			[][]int{
				[]int{1, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			0,
		},
	}

	for _, c := range cases {
		out := uniquePathsWithObstacles(c.grid)
		if out != c.n {
			t.Fatalf("grid: %v\nexpect path: %d, but got: %d\n", c.grid, c.n, out)
		}
	}
}
