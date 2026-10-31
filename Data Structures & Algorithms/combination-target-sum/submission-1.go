/*
2 5 6 9 t = 9
res cur = []
start = 0
remain = 9
i = 0
cur = [2]
start1 = 0
remain - 2 = 7
i = 0
cur = [2 2]
start2 = 0
remain - 2 = 5
*/

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
