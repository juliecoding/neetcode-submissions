func isPalindrome(s string) bool {
	lowercased := strings.ToLower(s)
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	normalized := reg.ReplaceAllString(lowercased, "")

	for i := range normalized {
		beginning := normalized[i]
		end := normalized[len(normalized) - 1 - i]

		if beginning != end {
			return false
		}
	}

	return true
}
