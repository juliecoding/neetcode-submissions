func longestConsecutive(nums []int) int {
	// Empty structs will take 0 bytes of memory, while booleans will take 1
	seen := make(map[int]struct{})
	for _, n := range nums {
		seen[n] = struct{}{}
	}

	max := 0
	for num, _ := range seen {
		// If num-1 is NOT found, this is the potential start of a sequence
		if _, ok := seen[num - 1]; !ok {
			count := 1
			
			// Using a 'while' loop so we can control iteration
			for {
				// Next consecutive number is found
				if _, exists := seen[num + count]; exists {
					count++
				} else {
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
