func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var storage [26]int

	// Идем по байтам строк одновременно
	for i := 0; i < len(s); i++ {
		storage[s[i]-'a']++
		storage[t[i]-'a']--
	}

	for _, count := range storage {
		if count != 0 {
			return false
		}
	}

	return true
}