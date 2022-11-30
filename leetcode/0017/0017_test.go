package _017

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
var mapping = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var result []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result = make([]string, 0)
	helper("", digits)
	return result
}

func helper(combination, digits string) {
	if len(digits) == 0 {
		result = append(result, combination)
	} else {
		digit := string(digits[0])
		letters := mapping[digit]
		for i := 0; i < len(letters); i++ {
			helper(combination+string(letters[i]), digits[1:])
		}
	}
}

func TestName(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}
