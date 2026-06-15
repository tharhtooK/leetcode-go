package prefixsum

import "testing"

func TestKSumSubarray(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{-1, -1, 1},
			0,
			1,
		},
		{
			[]int{1, 2},
			4,
			0,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
		{
			[]int{1, 1, 1},
			2,
			2,
		},
		{
			[]int{1, 2, -1},
			3,
			1,
		},
		{
			[]int{1, 2, -1, 1, 2},
			3,
			3,
		},
	}
	for _, test := range tests {
		result := KSumSubarray(test.nums, test.k)
		if result != test.expected {
			t.Errorf("expected %v got %v", test.expected, result)
		}
	}
}
