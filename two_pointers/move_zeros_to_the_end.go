package twopointers

func MoveZerosToTheEnd(nums []int) []int {
	/*
		Time Complexity: O(n)
		Space Complexity: o(1)
	*/
	first := 0
	for second := range len(nums) {
		if nums[second] != 0 {
			nums[first], nums[second] = nums[second], nums[first]
			first += 1
		}

	}
	return nums
}
