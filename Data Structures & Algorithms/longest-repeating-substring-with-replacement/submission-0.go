func characterReplacement(s string, k int) int {
	counts := [26]int{}

	left := 0
	res := 0
	maxf := 0

	for right, char := range s {
		counts[char-'A']++

		if counts[char-'A'] > maxf {
			maxf = counts[char-'A']
		}

		if (right-left + 1) - maxf > k {
			counts[s[left]-'A']--
			left++
		} 

		if (right - left + 1) > res {
			res = right-left+1
		}
	}
	return res
}
