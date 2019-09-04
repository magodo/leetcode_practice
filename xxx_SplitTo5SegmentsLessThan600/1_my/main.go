package main

import (
	"fmt"
	"strconv"
)

func main() {
	r := findSlice("123622349237", 5)
	fmt.Println(r)
	fmt.Println(len(r))
}

func findSlice(arr string, n int) [][]string {
	if n == 1 {
		num, _ := strconv.Atoi(arr)
		if num >= 600 {
			return nil
		}
		return [][]string{[]string{arr}}
	}

	result := [][]string{}
	arrLen := len(arr)
	if arrLen >= 1 {
		slices := findSlice(arr[:arrLen-1], n-1)
		if slices != nil {
			cslices := constructResult(slices, arr[arrLen-1:])
			result = append(result, cslices...)
		}
	}

	if arrLen >= 2 {
		slices := findSlice(arr[:arrLen-2], n-1)
		if slices != nil {
			cslices := constructResult(slices, arr[arrLen-2:])
			result = append(result, cslices...)
		}
	}

	if arrLen >= 3 {
		num, _ := strconv.Atoi(arr[arrLen-3:])
		if num <= 600 {
			slices := findSlice(arr[:arrLen-3], n-1)
			if slices != nil {
				cslices := constructResult(slices, arr[arrLen-3:])
				result = append(result, cslices...)
			}
		}
	}
	return result
}

func constructResult(result [][]string, substring string) [][]string {
	for idx := range result {
		result[idx] = append(result[idx], substring)
	}
	return result
}
