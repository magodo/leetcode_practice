package main

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	ttable := []struct {
		s string
		n int
	}{
		{"abc", 3},
		{"", 0},
		{"asdfssdsfk", 4},
		{"aabc", 3},
		{"dvdf", 3},
		{"abcabcbb", 3},
	}

	for _, test := range ttable {
		t.Logf("\nstring: %s\nn: %d\n", test.s, test.n)
		n := lengthOfLongestSubstring(test.s)
		if n != test.n {
			t.Fatalf("\nunexpected n: %d\n", n)
		}
	}
}
