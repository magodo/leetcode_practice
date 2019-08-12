package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length == 1 {
		_, _, stripped := stripK(nums1, nums2, 1)
		return float64(stripped[0])
	}
	n := length / 2
	k := n / 2
	for k > 0 {
		nums1, nums2, _ = stripK(nums1, nums2, k)
		n = n - k
		k = n / 2
	}

	nums1, nums2, stripped := stripK(nums1, nums2, 1)
	i := stripped[0]
	nums1, nums2, stripped = stripK(nums1, nums2, 1)
	j := stripped[0]
	if length%2 != 0 {
		return float64(j)
	}
	return float64(i+j) / 2
}

func stripK(nums1 []int, nums2 []int, k int) (new1, new2, stripped []int) {
	switch {
	case len(nums1) >= k && len(nums2) >= k:
		if nums1[k-1] > nums2[k-1] {
			return nums1, nums2[k:], nums2[:k]
		}
		return nums1[k:], nums2, nums1[:k]
	case len(nums1) >= k:
		return nums1[k:], nums2, nums1[:k]
	case len(nums2) >= k:
		return nums1, nums2[k:], nums2[:k]
	default:
		return nums1, nums2, []int{}
	}
}
