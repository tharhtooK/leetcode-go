package twopointers

import "testing"

func TestLargestContainer(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			[]int{2, 7, 8, 3, 7, 6},
			24,
		},
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 1},
			1,
		},
	}
	for _, test := range tests {
		result := LargestContainer(test.height)
		if result != test.expected {
			t.Errorf("expected %v got %v", test.expected, result)
		}
	}
}
