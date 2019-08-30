package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLIS(nums []int) int {
	maxLength := 0
	maxLengthEndWithN := map[int]int{}
	for _, n := range nums {
		maxLengthEndWithN[n] = 1
		for k := range maxLengthEndWithN {
			if k < n {
				maxLengthEndWithN[n] = max(maxLengthEndWithN[k]+1, maxLengthEndWithN[n])
			}
		}
		maxLength = max(maxLength, maxLengthEndWithN[n])
	}
	return maxLength
}
