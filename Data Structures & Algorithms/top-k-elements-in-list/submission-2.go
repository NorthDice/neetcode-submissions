func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _,v := range nums {
		m[v]++
	}
	arr := make([][]int, len(nums)+1)
	for key , v := range m {
		arr[v] = append(arr[v],key)
	}
	result := make([]int,0,k)
	
	for i := len(arr)-1; i >=0;i--{
		for _,num := range arr[i] {
			result = append(result, num)
			if len(result) == k {
            return result
        }
		}
	}
	return nil
}
