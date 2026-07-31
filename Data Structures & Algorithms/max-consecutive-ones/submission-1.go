func findMaxConsecutiveOnes(nums []int) int {
	var max, temp int
	for _, n := range nums {
		if n > 0 {
			temp += 1			
		} else {
			temp = 0			
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
