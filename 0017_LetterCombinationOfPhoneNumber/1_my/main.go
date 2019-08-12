package main

import "strings"

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var out []string
	for _, d := range digits {
		out = addNumber(out, string(d))
	}
	return out
}

func addNumber(currentCombines []string, digit string) (newCombinations []string) {
	if currentCombines == nil {
		return strings.Split(phoneMap[digit], "")
	}

	for _, c := range phoneMap[digit] {
		for _, currentCombine := range currentCombines {
			newCombinations = append(newCombinations, currentCombine+string(c))
		}
	}
	return newCombinations
}
