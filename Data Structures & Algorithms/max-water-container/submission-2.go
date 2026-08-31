func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	max := 0

	for left < right {
		height := calculateMin(heights[left], heights[right])
		width := right - left
		area := height * width

		if area > max {
			max = area
		}

		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func calculateMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
