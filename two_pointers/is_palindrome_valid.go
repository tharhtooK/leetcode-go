package twopointers

import (
	"strings"
	"unicode"
)

func IsAlNum(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return false
	}
	return true
}

func PalindromeValid(s string) bool {
	input := []rune(strings.ToLower(s))
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !IsAlNum(input[left]) {
			left += 1
		}
		for left < right && !IsAlNum(input[right]) {
			right -= 1
		}
		if input[left] != input[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
