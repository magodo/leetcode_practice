package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	xx := 0
	for y := x; y > 0; y /= 10 {
		xx = 10*xx + y%10
	}
	return xx == x
}
