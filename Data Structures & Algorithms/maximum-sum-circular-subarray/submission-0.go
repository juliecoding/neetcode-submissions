func maxSubarraySumCircular(nums []int) int {
	// Track the maximum subarray sum, the minimum subarray sum, and the total sum.
	// We need to track the total so we can know the potential wrapped total.
	// The wrapped subarray sum can be computed as:
	// total sum - minimum subarray sum
	maxSum := nums[0]
	maxAcc := 0

	minSum := nums[0]
	minAcc := 0

	totalSum := 0

	for _, n := range nums {
		// Kadane's algo for maximum non-empty subarray sum
		if maxAcc < 0 {
			maxAcc = 0
		}
		maxAcc += n
		if maxAcc > maxSum {
			maxSum = maxAcc
		}

		// Same approach but for minimum subarray sum
		if minAcc > 0 {
			minAcc = 0
		}
		minAcc += n
		if minAcc < minSum {
			minSum = minAcc
		}

		// Also track the total in order to calculate the wrapping case
		totalSum += n
	}

	// If all values are negative, return the largest single value.
	// We do this before calculating circularSum because if we don't, totalSum - minSum will be 0. It's greater than maxSum, but we're not allowed to return an empty subArray, so we take the largest value here instead.  
	if maxSum < 0 {
		return maxSum 
	}

	// Remove the smallest middle part to calculate largest possible wraparound without it
	circularSum := totalSum - minSum

	if circularSum > maxSum {
		return circularSum
	}

	return maxSum
}
