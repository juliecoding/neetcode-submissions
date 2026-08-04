func longestConsecutive(nums []int) int {
	max := 0
	seen := make(map[int]bool)

	for _, n := range nums {
		seen[n] = true
	}

	for key, _ := range seen {
		// If key-1 is NOT found, this is the potential start of a sequence
		if _, ok := seen[key - 1]; !ok {
			count := 0
			for i := key + 1; true; i++ {
				count++

				// Next consecutive number was not found
				if _, ok := seen[i]; !ok {
					break
				}
			}
			if count > max {
				max = count
			}
		}
	}

	return max
}
