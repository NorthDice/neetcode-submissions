func twoSum(nums []int, target int) []int {
    m:= make(map[int]int)

    for index, value := range nums {
        complement := target - value
        if e,ok := m[complement]; ok {
            return []int{e,index} 
        }

        m[value] = index
    }
    return nil
}
