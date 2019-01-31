package main

/**
* Space complexity: O(n). We store `nums` into an slice of structure `Element`
* Time complexity: O(n*lgn). We iterate the nums(O(n)), and in each iteration, use binary search (O(lgn)),
							 to search for the complement.
*/

import (
	"sort"
)

func twoSumArray(nums []int, target int) []int {
	elements := NewElements(nums)
	sort.Sort(elements)

	for i := 0; i < elements.Len()-1; i++ {
		searchCount := elements.Len() - i - 1
		offset := sort.Search(searchCount, func(index int) bool {
			return elements[i].Value+elements[index+i+1].Value >= target
		})

		if offset == searchCount {
			continue
		}

		// check whether the one at offset meets expectation
		if elements[i].Value+elements[i+offset+1].Value == target {
			return []int{elements[i].Index, elements[i+offset+1].Index}
		}
		// check whether the one before offset meets expectation
		if elements[i].Value+elements[i+offset].Value == target {
			return []int{elements[i].Index, elements[i+offset].Index}
		}
	}
	return nil
}

type Element struct {
	Value int
	Index int
}

type Elements []Element

func (e Elements) Len() int           { return len(e) }
func (e Elements) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Elements) Less(i, j int) bool { return e[i].Value < e[j].Value }

func NewElements(nums []int) Elements {
	var e Elements
	for i := range nums {
		e = append(e, Element{nums[i], i})
	}
	return e
}
