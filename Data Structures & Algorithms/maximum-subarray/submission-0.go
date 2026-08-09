func maxSubArray(nums []int) int {
    maxSum := nums[0]
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = 0
		}

		currentSum += nums[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}


	return maxSum
}
