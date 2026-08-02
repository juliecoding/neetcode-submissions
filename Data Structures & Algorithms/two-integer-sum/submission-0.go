func twoSum(nums []int, target int) []int {
	numStore := make(map[int]int)

	for i, n := range nums {
		complement := target - n

		if complementInd, ok := numStore[complement]; ok {
			return []int{complementInd, i}
		}
		numStore[n] = i
	}
	return []int{}
}
