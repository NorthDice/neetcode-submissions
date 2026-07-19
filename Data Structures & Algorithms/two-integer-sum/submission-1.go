func twoSum(nums []int, target int) []int {
    m:= make(map[int]int, len(nums))
    for i,v := range nums {
        find := target - v
        if e,ok := m[find];ok {
            return []int{e,i}
        }
        m[v] = i
    }
    return nil
}
