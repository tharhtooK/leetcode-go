package slidingwindow

import (
	"strings"
)

func LongestSubstringUniqueChars(s string) int {
	input := strings.Split(s, "")
	maxLen := 0
	hasSeen := map[string]int{}
	left, right := 0, 0
	for right < len(input) {
		index, ok := hasSeen[input[right]]
		for ok && index >= left {
			delete(hasSeen, input[left])
			left = index + 1
		}

		hasSeen[input[right]] = right
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}
