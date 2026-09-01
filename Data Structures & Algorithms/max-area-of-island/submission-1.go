func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 {
		return 0
	}
	area := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				currentArea := dfs(grid,i,j)
				area = max(area,currentArea)
			}
		}
	}
	return area
}


func dfs(grid [][]int, i,j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j]== 0 {
		return 0
	}

	grid[i][j] = 0

	area := 1
	area += dfs(grid,i-1,j)
	area += dfs(grid,i,j-1)
	area += dfs(grid,i+1,j)
	area += dfs(grid,i,j+1)

	return area
}