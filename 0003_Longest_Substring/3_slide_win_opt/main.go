// prefer v2
package main

func lengthOfLongestSubstring(s string) int {
	ss := []rune(s)
	n := len(ss)
	m := make(map[rune]int) // char -> indicies in 's'(but not used)
	maxlen := 0
	for i, j := 0, 0; j < n; {
		if _, ok := m[ss[j]]; !ok {
			// add char into map
			m[ss[j]] = j
			j++
		} else {
			// ss[j] has already existed in [i, j)
			// TRICKY: since we store the indicies of each char, we can lookup hashtable
			//         and skip the chars before and includes the char which is duplicated with
			//         the current one (i.e. ss[j]).
			//         the most tricky part is the condition below, this is because we no
			//         longer guarantee 'm' contain the chars in [i,j), rather, we are allowing
			//         'm' contains more, like [a,b] + [i,j). The fowllowing condition is used
			//         to avoid 'i' being updated when s[j] == s[a] (or s[b])
			if m[ss[j]]+1 > i {
				i = m[ss[j]] + 1
			}
			m[ss[j]] = j
			j++
		}
		if j-i > maxlen {
			maxlen = j - i
		}
	}
	return maxlen
}
