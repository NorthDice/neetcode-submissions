func combinationSum(nums []int, target int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start,remain int)
	backtrack = func(start,remain int) {
		if remain == 0 {
			combo := make([]int,len(current))
			copy(combo, current)
			result = append(result,combo)
			return
		}
		if remain < 0 {
			return
		}

		for i := start; i < len(nums); i++{
			current = append(current, nums[i])

			backtrack(i,remain-nums[i])

			current = current[:len(current)-1]
		}
	}
	backtrack(0,target)
	return result
	
}
