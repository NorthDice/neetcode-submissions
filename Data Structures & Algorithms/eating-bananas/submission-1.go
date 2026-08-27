func minEatingSpeed(piles []int, h int) int {
    left := 1
    right := 0 

    for _, v := range piles {
        right = max(right, v)
    }

    for left < right {
        mid := left + (right - left) / 2

        hoursSpent := 0
        for _, pile := range piles {
            hoursSpent += (pile + mid - 1) / mid
        }

        if hoursSpent <= h {
            right = mid 
        } else {
            left = mid + 1
        }
    }
    return left
}
