func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	for i := range grid {
		for j, val := range grid[i] {
			if val == 1 {
				maxArea = max(maxArea, dfs(grid, i, j))
			}
		}
	}

	return maxArea
}

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func dfs(grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0 
	area := 1

	for _, d := range directions {
		area += dfs(grid, r+d[0], c+d[1])
	}

	return area
}