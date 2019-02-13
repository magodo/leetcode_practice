package main

func longestPalindrome(s string) string {
	ss := []rune(s)

	var maxstr []rune

	if len(ss) == 1 {
		return s
	}

	for i := 0; i < len(ss)-1; i++ {
		substr := calcLen(ss, i, i)
		if len(substr) > len(maxstr) {
			maxstr = substr
		}

		substr = calcLen(ss, i, i+1)
		if len(substr) > len(maxstr) {
			maxstr = substr
		}
	}
	return string(maxstr)
}

func calcLen(s []rune, l, r int) []rune {
	for ; r < len(s) && l >= 0; l, r = l-1, r+1 {
		if s[l] != s[r] {
			break
		}
	}
	return s[l+1 : r]
}
