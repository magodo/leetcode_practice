package main

import (
	"math"
)

func reverse(x int) int {
	var remainder, y int
	for remainder, y = 0, 0; x != 0; {
		x, remainder = x/10, x%10
		if y > math.MaxInt32/10 || (y == math.MaxInt32/10 && remainder > math.MaxInt32%10) {
			return 0
		}
		if y < math.MinInt32/10 || (y == math.MinInt32/10 && remainder < math.MinInt32%10) {
			return 0
		}
		y = y*10 + remainder
	}
	return y
}
