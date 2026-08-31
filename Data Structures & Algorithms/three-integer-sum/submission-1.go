import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	triplets := make([][]int, 0)
	sum := 0

	// Only looping to length - 2 to allow room for j and k
	for i := 0; i < len(nums) - 2; i++ {
		// If nums[i] > 0, all following values are positive, so they can't equal 0
		if nums[i] > 0 {
			break
		}

		// Skip duplicate values
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				continue
			}
			if sum > 0 {
				k--
				continue
			}

			triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
			j++
			k--

			// If the next value(s) are duplicates, increase until they're not
			for j < k && nums[j] == nums[j-1] {
				j++
			}
			for j < k && nums[k] == nums[k+1] {
				k--
			}
		}
	}	
	return triplets
}
