func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]int)
	res := 0
	l := 0
	runes := []rune(s)
	for r, char := range runes {
		if pervIdx, ok := set[char]; ok && pervIdx >= l{
			l = pervIdx + 1
		}
		set[char] = r
		if r - l + 1 > res {
			res = r-l+1
		}
	}
	return res
}
