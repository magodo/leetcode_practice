package main

import (
	"math"
)

func findUnsortedSubarray(nums []int) int {
	lefts := []int{nums[0]}
	rights := []int{nums[len(nums)-1]}

	isUp := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			break
		}
		if nums[i] > nums[i-1] {
			isUp = true
		}
		lefts = append(lefts, nums[i])
	}
	if !isUp {
		lefts = []int{}
	}
	if len(lefts) == len(nums) {
		return 0
	}

	isUp = false
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			break
		}
		if nums[i] < nums[i+1] {
			isUp = true
		}
		rights = append([]int{nums[i]}, rights...)
	}
	if !isUp {
		rights = []int{}
	}

	min := int(math.MaxInt64)
	max := int(math.MinInt64)
	for i := 0; i < len(nums)-len(rights); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	for i := len(lefts); i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	if min == max && len(lefts) == 0 && len(rights) == 0 {
		return 0
	}

	var start, end int
	for start = len(lefts) - 1; start >= 0; start-- {
		if lefts[start] <= min {
			break
		}
	}

	for end = 0; end < len(rights); end++ {
		if rights[end] >= max {
			break
		}
	}
	end += len(nums) - len(rights)
	return end - start - 1
}
