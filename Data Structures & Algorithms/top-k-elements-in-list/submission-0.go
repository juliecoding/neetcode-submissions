func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	// The `+ 1` comes from the fact that we want the counts to match the indexes, and they'll be off by one if we don't. (E.g., if the `nums` slice contains 5 total elements and they're all the same, we want to store that count at index 5, not index 4.) 
	bucket := make([][]int, len(nums) + 1)

	for _, n := range nums {
		counts[n]++
	}

	for num, count :=  range counts {
		bucket[count] = append(bucket[count], num)
	}

	// preallocate
	output := make([]int, 0, k)

	// Why this isn't O(n^2): 
	// Each distinct number from `nums` appears in exactly one bucket. So the total number of iterations across all inner loops is equal to the number of distinct numbers, which is at most n.
	for i := len(bucket) - 1; i > 0 && len(output) < k; i-- {
		for _, num := range bucket[i] {
			output = append(output, num)
			
			if len(output) == k {
				return output
			}
		}
	}

	return output
}
