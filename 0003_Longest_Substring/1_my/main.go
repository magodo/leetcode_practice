package main

import "errors"

type Set struct {
	m map[rune]int
}

type Element struct {
	value rune
	index int
}

func (s *Set) Add(e *Element) error {
	if _, ok := s.m[e.value]; ok {
		return errors.New("exists")
	}
	s.m[e.value] = e.index
	return nil
}
func (s *Set) Replace(e *Element) {
	collide := &Element{e.value, s.m[e.value]}
	for value, index := range s.m {
		if index <= collide.index {
			delete(s.m, value)
		}
	}
	s.Add(e)
}

func (s *Set) Size() int {
	return len(s.m)
}

func NewSet() *Set {
	return &Set{m: make(map[rune]int)}
}

func lengthOfLongestSubstring(s string) int {
	set := NewSet()
	maxLen := 0
	for index, value := range s {
		e := &Element{value, index}
		if set.Add(e) == nil {
			continue
		}

		// duplication occurs
		if set.Size() > maxLen {
			maxLen = set.Size()
		}
		set.Replace(e)
	}

	if set.Size() > maxLen {
		maxLen = set.Size()
	}

	return maxLen
}
