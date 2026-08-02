func groupAnagrams(strs []string) [][]string {
	// Arrays can be map keys and can be compared for equivalency due to their fixed nature
	m := make(map[[26]int][]string)

	for _, s := range strs {
		signature := [26]int{}
		for _, t := range s {

			// Get index of the letter in the English alphabet by subtracting 'a' (which has ASCII code 97)
			indexInAlphabet := t - 'a'
			signature[indexInAlphabet]++
		}

		m[signature] = append(m[signature], s)
	}

	output := make([][]string, 0)
	for _, value := range m {
		output = append(output, value)
	}

	return output
}
