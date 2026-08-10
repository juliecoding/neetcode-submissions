func getConcatenation(nums []int) []int {
	// This is potentially more efficient than using `append`
	// Because we know we won't reallocate memory.
	ans := make([]int, len(nums) * 2)

	// Copy can be faster than a loop when working with large slices
	copy(ans, nums)
	copy(ans[len(nums):], nums)

	return ans
}
