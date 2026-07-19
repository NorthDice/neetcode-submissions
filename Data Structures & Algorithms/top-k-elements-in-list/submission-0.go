func topKFrequent(nums []int, k int) []int {
   m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }

    buckets := make([][]int, len(nums)+1)
    for num, freq := range m {
        buckets[freq] = append(buckets[freq], num)
    }

    result := make([]int, 0, k) // Оптимизация: сразу задаем емкость k
    for i := len(buckets) - 1; i > 0; i-- {
        for _, num := range buckets[i] {
            result = append(result, num)
            if len(result) == k {
                return result
            }
        }
    }

    return result

}
