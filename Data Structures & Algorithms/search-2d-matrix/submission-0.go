func searchMatrix(matrix [][]int, target int) bool {
	isTargetInArray := false
	for _,arr:= range matrix {
		isTargetInArray = binarySearch(arr,target)
		if isTargetInArray == true {
			return true
		}
	}
	return isTargetInArray
}

func binarySearch(arr []int,target int) bool {
	l,r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}