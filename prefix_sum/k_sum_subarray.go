package prefixsum

func KSumSubarray(nums []int, k int) int {
	var count int = 0
	var currPrefixSum int = 0
	sumFreqMap := map[int]int{0: 1}

	for _, num := range nums {
		currPrefixSum += num
		if freq, ok := sumFreqMap[currPrefixSum-k]; ok {
			count += freq
		}
		sumFreqMap[currPrefixSum] += 1
	}
	return count
}
