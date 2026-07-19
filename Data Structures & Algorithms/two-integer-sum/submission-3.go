func twoSum(nums []int, target int) []int {
    stored := make(map[int]int)
    for i,v := range nums {
        f := target - v
        if _,ok := stored[f]; !ok {
            stored[v] = i
        } else {
            return []int{stored[f], i}
        }
    }
    return []int{}
}
