func maxProfit(prices []int) int {
	l := 0
	r := 1
	max := 0
	for r < len(prices){
		if prices[l] > prices[r] {
			l = r
		} else {
			if prices[r] - prices[l] > max {
				max = prices[r]- prices[l]
			}
		}
		r++
	}
	

	return max
}
