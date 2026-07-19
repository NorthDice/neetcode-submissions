func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{}, len(nums))

    for _, v := range nums {
        if _, exists := m[v]; exists {
            return true
        }
        m[v] = struct{}{}
    }
    return false
}
