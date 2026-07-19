func twoSum(nums []int, target int) []int {
    m:= make(map[int]int)
    //  Идем по масиву
    for i,v := range nums {
        find := target - v
        if e,ok := m[find];ok {
            return []int{e,i}
        }
        m[v] = i
    }
    return nil
    // Отнимать от таргета число
    // Искать есть ли это число в мапе
    // Если да вернуть индекс этого числа и число которое сейчас
    // ИНаче записать число и индекс
}
