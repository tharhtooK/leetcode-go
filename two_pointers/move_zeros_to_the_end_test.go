package twopointers

import (
	"slices"
	"testing"
)

func TestMoveZerosToTheEnd(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{0, 1, 0, 3, 2},
			[]int{1, 3, 2, 0, 0},
		},
	}
	for _, test := range tests {
		result := MoveZerosToTheEnd(test.nums)
		if !slices.Equal(result, test.expected) {
			t.Errorf("\n\n=====> expected %v got %v", test.expected, result)
		}
	}
}
