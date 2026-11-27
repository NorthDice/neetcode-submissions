func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	unreachable := amount + 1
	for i := 1; i <= amount; i++ {
		dp[i] = unreachable
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == unreachable {
		return -1
	}

	return dp[amount]
}
