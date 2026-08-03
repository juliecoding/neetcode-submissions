type Solution struct{}

const DELIMITER = "-"

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, str := range strs {
		res = res + strconv.Itoa(len(str)) + DELIMITER + str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}

	i := 0 

	// while-style loop for control of incrementation
	for i < len(encoded) {
		j := i
		// Alternatively, encoded[j] != DELIMITER[0], which is a little more efficient
		for string(encoded[j]) != DELIMITER {
			j++
		}

		numberOfChars, err := strconv.Atoi(encoded[i:j])
		if err != nil {
			return res
		}

		// Add one so we don't include the delimiter
		start := j + 1
		end := start + numberOfChars

		res = append(res, encoded[start:end])

		i = end
	}

	return res
}