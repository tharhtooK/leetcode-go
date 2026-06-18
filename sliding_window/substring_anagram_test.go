package slidingwindow

import (
	"testing"
)

func TestSubstringAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected int
	}{
		{
			"caabab",
			"aba",
			2,
		},
		{
			"cbaebabacd",
			"abc",
			2,
		},
		{
			"abab",
			"ab",
			3,
		},
	}
	for _, test := range tests {
		result := SubstringAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("expected %v got %v", test.expected, result)
		}
	}
}
