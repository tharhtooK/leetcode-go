package slidingwindow

import (
	"testing"
)

func TestLongestSubstringUniqueChars(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			"abcba",
			3,
		},
		{
			"cabcdeca",
			5,
		},
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}
	for _, test := range tests {
		result := LongestSubstringUniqueChars(test.s)
		if result != test.expected {
			t.Errorf("expected %v got %v", test.expected, result)
		}
	}
}
