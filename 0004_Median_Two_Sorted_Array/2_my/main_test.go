package main

import "testing"

func isSliceEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestMergeSliceMinN(t *testing.T) {
	cases := []struct {
		s1 []int
		s2 []int
		n  int
		s  []int
	}{
		{
			[]int{1, 5, 7, 8},
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 5},
		},
		{
			[]int{1, 2, 3},
			[]int{},
			2,
			[]int{1, 2},
		},
		{
			[]int{},
			[]int{},
			3,
			[]int{},
		},
	}

	for _, c := range cases {
		out := mergeSliceForMinN(c.s1, c.s2, c.n)
		if !isSliceEqual(out, c.s) {
			t.Fatalf("\nout: %v\nexpect: %v\n", out, c.s)
		}
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		s1 []int
		s2 []int
		m  float64
	}{
		{
			[]int{1},
			[]int{},
			1.0,
		},
		{
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			[]int{1, 3, 4},
			[]int{2},
			2.5,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{},
			2.5,
		},
		{
			[]int{1, 3, 8, 12},
			[]int{5, 6, 7, 8, 9},
			7,
		},
		{
			[]int{},
			[]int{1, 2, 3, 4, 5, 6},
			3.5,
		},
	}

	for _, c := range cases {
		median := findMedianSortedArrays(c.s1, c.s2)
		if median != c.m {
			t.Fatalf("\ns1: %v\ns2: %v\nmedian: %v (expected: %v)\n", c.s1, c.s2, median, c.m)
		}
	}
}

func mergeSliceForMinN(s1, s2 []int, n int) []int {
	s := []int{}
loop:
	for i := 0; i < n; i++ {
		switch {
		case len(s1) == 0:
			var idx int
			if len(s2) > n-i {
				idx = n - i
			} else {
				idx = len(s2)
			}
			s = append(s, s2[:idx]...)
			break loop
		case len(s2) == 0:
			var idx int
			if len(s1) > n-i {
				idx = n - i
			} else {
				idx = len(s1)
			}
			s = append(s, s1[:idx]...)
			break loop
		default:
			if s1[0] > s2[0] {
				s = append(s, s2[0])
				s2 = s2[1:]
			} else {
				s = append(s, s1[0])
				s1 = s1[1:]
			}
		}
	}
	return s
}
