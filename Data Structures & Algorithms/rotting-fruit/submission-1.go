type point struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	queue := make([]point, 0)
	fresh := 0

	for r := range rows {
		for c := range cols {
			switch grid[r][c] {
			case 2:
				queue = append(queue, point{row: r, col: c})
			case 1:
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	tryRot := func(nr, nc int) {
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			return
		}
		if grid[nr][nc] != 1 {
			return
		}

		grid[nr][nc] = 2
		fresh--
		queue = append(queue, point{row: nr, col: nc})
	}

	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		levelSize := len(queue)
		minutes++

		for range levelSize {
			curr := queue[0]
			queue = queue[1:]

			tryRot(curr.row-1, curr.col) 
			tryRot(curr.row+1, curr.col) 
			tryRot(curr.row, curr.col-1) 
			tryRot(curr.row, curr.col+1) 
		}
	}

	if fresh > 0 {
		return -1
	}

	return minutes
}