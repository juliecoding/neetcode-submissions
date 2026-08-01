func replaceElements(arr []int) []int {
	rightMax := -1 
	for i := len(arr) - 1; i >= 0; i-- {
		current := arr[i]
		arr[i] = rightMax

		if current > rightMax {
			rightMax = current
		}
	}

	return arr
}
