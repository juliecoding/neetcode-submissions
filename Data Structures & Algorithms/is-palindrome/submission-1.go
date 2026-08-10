func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	// Only loop halfway through the string
	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphaNum(s[left]) {
			left++
		}

		for left < right && !isAlphaNum(s[right]) {
			right--
		}

		// You could use strings.ToLower() here, but by treating these as bytes, you avoid the extra work of string conversion and allocations.
		// (The instructions specify the input is guaranteed to be printable ASCII.)
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}


	return true
}

func isAlphaNum(c byte) bool {
	return ('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9')
}

func toLower(c byte) byte {
	if 'A' <= c && c <='Z' {
		return c + ('a' - 'A')
	}
	return c
}