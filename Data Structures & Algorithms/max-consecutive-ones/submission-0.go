func findMaxConsecutiveOnes(nums []int) int {
	var max, temp int
	for _, n := range nums {
		if n > 0 {
			temp += 1			
			if temp > max {
				max = temp
			}
		} else {
			temp = 0			
		}
	}
	return max
}
