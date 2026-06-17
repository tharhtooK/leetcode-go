package twopointers

import "testing"

func TestPalindromeValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"racecar",
			true,
		},
		{
			"a dog! a panic in a pagoda.",
			true,
		},
	}
	for _, test := range tests {
		result := PalindromeValid(test.s)
		if result != test.expected {
			t.Errorf("expected %v got %v", test.expected, result)
		}
	}
}
