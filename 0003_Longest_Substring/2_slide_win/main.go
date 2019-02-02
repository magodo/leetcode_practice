package main

func lengthOfLongestSubstring(s string) int {
	ss := []rune(s)
	n := len(ss)
	m := make(map[rune]int) // char -> indicies in 's'(but not used)
	maxlen := 0
	for i, j := 0, 0; j < n && i < n; {
		if _, ok := m[ss[j]]; !ok {
			// add char into map
			m[ss[j]] = j
			j++
		} else {
			// ss[j] has already existed in [i, j)
			delete(m, ss[i])
			i++
		}
		if j-i > maxlen {
			maxlen = j - i
		}
	}
	return maxlen
}
