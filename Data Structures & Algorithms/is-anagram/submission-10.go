func isAnagram(s string, t string) bool {
	var storage [26]int
	if len(s) != len(t) {
		return false
	}

	for _,v := range s {
		storage[v-'a']++
	}
	for _,v := range t {
		storage[v-'a']--
	}
	for _,v := range storage {
		if v != 0 {
			return false
		}
	}
	return true

}
