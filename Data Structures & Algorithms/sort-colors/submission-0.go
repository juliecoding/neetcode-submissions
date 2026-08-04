func sortColors(nums []int) {
	counts := []int{0, 0, 0}

	for _, n := range nums {
		counts[n]++
	}

	cursor := 0
	for i := range counts {
		for _ = range counts[i] {
			// The value IS the index in counts
			nums[cursor] = i
			cursor++
		}
   }
}
