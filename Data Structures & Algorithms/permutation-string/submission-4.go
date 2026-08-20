func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var count1, count2 [26]int

	for i,v := range s1 {
		count1[v-'a']++
		count2[s2[i]-'a']++
	}

	if count1 == count2 {
		return true
	}

	for i,ch := range s2[len(s1):] {
		count2[ch-'a']++

		count2[s2[i]-'a']--

		if count1 == count2 {
			return true
		}
	}

	return false
}
