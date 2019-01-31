package main

func twoSumOnePassHash(nums []int, target int) []int {
	m := make(map[int]int) // [value]: indices
	for i := range nums {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{i, m[complement]}
		}
		m[nums[i]] = i
	}
	return nil
}
