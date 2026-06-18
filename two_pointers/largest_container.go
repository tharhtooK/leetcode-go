package twopointers

func LargestContainer(height []int) int {
	max_water := 0
	left, right := 0, len(height)-1
	for left < right {
		water := min(height[left], height[right]) * (right - left)
		max_water = max(max_water, water)
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return max_water
}
