package main

func Reverse(s []rune, idx int) {
	s1, s2 := s[:idx], s[idx:]
	reverse(s1)
	reverse(s2)
	reverse(s)
}

func reverse(s []rune) {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		c := s[start]
		s[start] = s[end]
		s[end] = c
	}
}
