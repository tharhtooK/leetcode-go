package prefixsum

func ProductArrayWithoutSelf(nums []int) []int {
	/*
		Time Complexity: O(n^2)
		Space Complexity: o(n)
	*/
	products := []int{1}
	for i := 1; i < len(nums); i++ {
		products = append(products, products[i-1]*nums[i-1])
	}

	right_product := 1
	for i := len(nums) - 2; i > -1; i-- {
		right_product *= nums[i+1]
		products[i] = right_product * products[i]
	}
	return products
}
