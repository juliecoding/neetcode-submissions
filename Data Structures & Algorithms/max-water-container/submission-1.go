func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxArea := 0

	for left < right {
		var min int
		if heights[left] <= heights[right] {
			min = heights[left]
		} else {
			min = heights[right]
		}

		area := (right - left) * min
		if area > maxArea {
			maxArea = area
		}

		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
