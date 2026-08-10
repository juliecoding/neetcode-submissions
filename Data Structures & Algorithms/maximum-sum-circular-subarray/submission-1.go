func maxSubarraySumCircular(nums []int) int {
	maxSum := nums[0]
	maxAcc := 0

	minSum := nums[0]
	minAcc := 0

	totalSum := 0

	for _, n := range nums {
		// Kadane's for max and min subarrays 
		maxAcc = max(n, maxAcc + n)
		maxSum = max(maxSum, maxAcc)

		minAcc = min(n, minAcc + n)
		minSum = min(minSum, minAcc)

		totalSum += n
	}

	// If all values are negative, return the largest single value
	// Do this before calculating circularSum because the circularSum would be 0 in this case (representing 0 values from the array), and we're not allowed to return an empty sub-array
	if maxSum < 0 {
		return maxSum
	}

	// Take out the smallest "middle" subArray
	circularSum := totalSum - minSum

	return max(maxSum, circularSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}