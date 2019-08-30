package main

func uniquePaths(m int, n int) int {
	x := m
	if m > n {
		x = n
	}
	return combination(m+n-2, x-1)
}

// combination calculate how many combinations of b objects from total a objects
func combination(a, b int) int {
	if a == 0 || b == 0 || a == b {
		return 1
	}
	divider := 1
	for i := a; i > a-b; i-- {
		divider *= i
	}
	dividee := 1
	for i := b; i > 0; i-- {
		dividee *= i
	}

	return divider / dividee
}
