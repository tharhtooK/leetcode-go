package prefixsum

import (
	"slices"
	"testing"
)

func TestProductArrayWithoutSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{2, 3, 1, 4, 5},
			[]int{60, 40, 120, 30, 24},
		},
	}
	for _, test := range tests {
		result := ProductArrayWithoutSelf(test.nums)
		if !slices.Equal(result, test.expected) {
			t.Errorf("\n\n=====> expected %v got %v", test.expected, result)
		}
	}
}
