package main

func IsSubstring(sa, sb string) bool {
	return (alphsToBitPattern(sa) | alphsToBitPattern(sb)) == alphsToBitPattern(sa)
}

func alphsToBitPattern(s string) uint32 {
	n := uint32(0)
	for _, c := range s {
		n |= 1 << uint32(c-'A')
	}
	return n
}
