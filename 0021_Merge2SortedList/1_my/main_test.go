package main

import (
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		l1 []int
		l2 []int
		lm []int
	}{
		{
			[]int{1, 2, 4},
			[]int{1, 3, 4},
			[]int{1, 1, 2, 3, 4, 4},
		},
		{
			[]int{},
			[]int{1, 3, 4},
			[]int{1, 3, 4},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
	}

	for _, c := range cases {
		// construct list
		l1Head := FromSlice(c.l1)
		l2Head := FromSlice(c.l2)
		out := mergeTwoLists(l1Head, l2Head).ToSlice()
		if !reflect.DeepEqual(out, c.lm) {
			t.Fatalf("not equal:\noutput: %v\nexpected: %v\n", out, c.lm)
		}
	}
}
