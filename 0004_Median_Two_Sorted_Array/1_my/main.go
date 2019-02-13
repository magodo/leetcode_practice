package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	tmpIdx := (len1 + len2) / 2

	if (len1+len2)%2 == 1 {
		tmpIdx = tmpIdx + 1
	}
	cIdx := tmpIdx / 2

	for cIdx > 0 {

		switch {
		case len(nums1) < cIdx:
			nums2 = nums2[cIdx:]
		case len(nums2) < cIdx:
			nums1 = nums1[cIdx:]
		default:
			if nums1[cIdx-1] < nums2[cIdx-1] {
				nums1 = nums1[cIdx:]
			} else {
				nums2 = nums2[cIdx:]
			}
		}

		tmpIdx = tmpIdx - cIdx
		cIdx = tmpIdx / 2
	}

	s := mergeSliceForMinN(nums1, nums2, 2)
	if (len1+len2)%2 == 0 {
		return float64(s[0]+s[1]) / 2

	}
	return float64(s[0])
}

func mergeSliceForMinN(s1, s2 []int, n int) []int {
	s := []int{}
loop:
	for i := 0; i < n; i++ {
		switch {
		case len(s1) == 0:
			var idx int
			if len(s2) > n-i {
				idx = n - i
			} else {
				idx = len(s2)
			}
			s = append(s, s2[:idx]...)
			break loop
		case len(s2) == 0:
			var idx int
			if len(s1) > n-i {
				idx = n - i
			} else {
				idx = len(s1)
			}
			s = append(s, s1[:idx]...)
			break loop
		default:
			if s1[0] > s2[0] {
				s = append(s, s2[0])
				s2 = s2[1:]
			} else {
				s = append(s, s1[0])
				s1 = s1[1:]
			}
		}
	}
	return s
}
