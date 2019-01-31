package main

import "testing"

func TestSearch1(t *testing.T) {

	tables := []struct {
		nums     []int
		target   int
		indicies []int
	}{
		{[]int{3, 5, 7, 2, 9}, 10, []int{0, 2}},
		{[]int{3, 5, 7, -1, 2, 9}, 6, []int{2, 3}},
		{[]int{3, 5, 7, -1, 2, 9}, 600, nil},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	funcs := []struct {
		f    func([]int, int) []int
		desc string
	}{
		{twoSumArray, "array sort"},
		{twoSumOnePassHash, "one pass hash map"},
	}

	for _, f := range funcs {
		t.Logf("\n\nMethod: %s\n\n", f.desc)
		for _, table := range tables {
			indicies := f.f(table.nums, table.target)
			t.Logf("\ntarget: %v\nnums: %v\nindicies: %v", table.target, table.nums, indicies)
			if table.indicies == nil {
				if indicies != nil {
					t.Errorf("should be no solution")
				} else {
					continue
				}
			}
			if len(indicies) != len(table.indicies) {
				t.Errorf("indicies amount incorrect")
			}
			if !((indicies[0] == table.indicies[0] && indicies[1] == table.indicies[1]) ||
				(indicies[0] == table.indicies[1] && indicies[1] == table.indicies[0])) {
				t.Errorf("failed")
			}
		}
	}
}
