package main

import (
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	cases := []struct {
		input  string
		expect []string
	}{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"2",
			[]string{"a", "b", "c"},
		},
	}

	for _, c := range cases {
		out := letterCombinations(c.input)
		if len(c.expect) != len(out) {
			t.Fatalf("expect: %v, but got %v", c.expect, out)
		}
	loop:
		for _, substring1 := range c.expect {
			for _, substring2 := range out {
				if substring1 == substring2 {
					continue loop
				}
			}
			t.Fatalf("%s combination not found (but contains %v)", substring1, out)
		}
		if reflect.DeepEqual(out, c.expect) {
		}
	}
}
