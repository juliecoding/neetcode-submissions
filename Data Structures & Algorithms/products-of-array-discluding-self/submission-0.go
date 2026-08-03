func productExceptSelf(nums []int) []int {
	prefixes := make([]int, len(nums))
	prefixes[0] = 1
	suffixes := make([]int, len(nums))
	suffixes[len(suffixes) - 1] = 1

	// example 
	// [1,2,4,6]
	for i := 1; i < len(nums); i++ {
		prefixes[i] = prefixes[i - 1] * nums[i - 1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffixes[i] = suffixes[i + 1] * nums[i + 1]
	}

	res := make([]int, len(nums))
	for i := range res {
		res[i] = prefixes[i] * suffixes[i]
	}

	return res
}
