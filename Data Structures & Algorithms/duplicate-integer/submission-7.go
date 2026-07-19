func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{}, len(nums))

    for _, v := range nums {
        if _, exist := m[v]; exist {
            return true
        }
        m[v] = struct{}{}
    }

    return false
}
