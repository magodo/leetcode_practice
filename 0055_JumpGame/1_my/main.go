package main

func canJump(nums []int) bool {
	for rRight, idx := 0, 0; nums[idx]+idx < len(nums)-1; {
		rLeft, rRight := rRight+1, nums[idx]+idx
		rMax := 0
		for i := range nums[rLeft : rRight+1] {
			i += rLeft
			jumpLength := nums[i] - (rRight - i)
			if jumpLength > rMax {
				rMax = jumpLength
				idx = i
			}
		}
		if rMax == 0 {
			return false
		}
	}
	return true
}
