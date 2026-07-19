func maxArea(heights []int) int {
	l,r := 0, len(heights) - 1
	res := 0
	for l < r {
		area := (r - l) * min(heights[l], heights[r])
		res = max(res,area) 
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
		
	}
	return res
}	
