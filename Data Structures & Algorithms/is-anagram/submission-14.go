func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var counts [26]int

    for _, r := range s {
        counts[r-'a']++
    }

    for _, r := range t {
        counts[r-'a']--
    }

    for _,v := range counts {
        if v != 0 {
            return false
        }
    }
    return true
}