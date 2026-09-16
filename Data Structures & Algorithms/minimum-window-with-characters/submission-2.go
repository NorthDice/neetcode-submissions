func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }

	var countsT [128]int
	required := 0
	for _, char := range t {
		if countsT[char] == 0 {
			required++
		}
		countsT[char]++
	}

	var windowCounts[128]int
	bestStart :=0
	minLen := len(s)+1
	l := 0
	have := 0

	for r, char := range s {
		windowCounts[char]++

		if countsT[char] > 0 && windowCounts[char] == countsT[char]{
			have++
		}

		for have == required {
			currentLen := r-l+1
			if currentLen < minLen {
				minLen = currentLen
				bestStart = l
			}

			leftChar := s[l]
			windowCounts[leftChar]--

			if countsT[leftChar] > 0 && windowCounts[leftChar] < countsT[leftChar] {
				have--
			}
			l++
		}
		
	}
	if minLen > len(s) {
		return ""
	}

	return s[bestStart:bestStart+minLen]
}
