func sortColors(nums []int) {
	counts := []int{0, 0, 0}

	for _, n := range nums {
		counts[n]++
	}

	cursor := 0
	for i := range counts {
		// This could also be `for count[i] > 0 {` or `for j := 0; j < counts[i]; j++`
		// I kind of like this because it shows I'm disinterested in the iterator's value
		for _ = range counts[i] {
			// The value IS the index, not the count
			nums[cursor] = i
			cursor++
		}
   }
}
