// 2,20,4,10,3,4,5 
func longestConsecutive(nums []int) int {
	ints := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		ints[num] = struct{}{}
	}

	
	maxSeq := 0

	for v := range ints {
		if _, ok := ints[v-1];ok {
			continue
		} 
		currentNum := v
        currentSeq := 1

		for {
			if _, ok := ints[currentNum+1]; !ok {
				break
			}
			currentNum++
			currentSeq++
		}
		
		maxSeq = max(maxSeq, currentSeq)
		
	}
	return maxSeq
}
