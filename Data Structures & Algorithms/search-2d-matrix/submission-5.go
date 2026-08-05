func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	l,r := 0, row * col -1

	for l<=r {
		mid := l+(r-l)/2
		midVal := matrix[mid/col][mid % col]
		if midVal == target {
			return true
		}
		if midVal > target {
			r = mid -1
		} else {
			l = mid + 1
		}
	}
	return false
}

