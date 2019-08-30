package main

func generateParenthesis(n int) []string {
	return backtrack(n, n)
}

func backtrack(op, ed int) []string {
	if op == 0 {
		s := ""
		for i := 0; i < ed; i++ {
			s += ")"
		}
		return []string{s}
	}
	s1 := backtrack(op-1, ed)
	for idx := range s1 {
		s1[idx] = "(" + s1[idx]
	}
	if op == ed {
		return s1
	}
	s2 := backtrack(op, ed-1)
	for idx := range s2 {
		s2[idx] = ")" + s2[idx]
	}
	return append(s1, s2...)
}
