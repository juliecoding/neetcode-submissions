func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[byte]int)
	countT := make(map[byte]int)

	for i := range s {
		countS[s[i]]++
		countT[t[i]]++
	}

	for k, v := range countS{
		if countT[k] != v {
			return false
		}
	}

	return true
}
