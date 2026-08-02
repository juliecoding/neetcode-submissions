func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, n := range nums {
		complement := target - n

		if complementIndex, ok := seen[complement]; ok {
			return []int{complementIndex, i}
		}
		seen[n] = i
	}
	return []int{}
}
