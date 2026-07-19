func isAnagram(s string, t string) bool {
      if len(s) != len(t) {
        return false
    }
    
    m := make(map[rune]int, len(s))
    z := make(map[rune]int, len(t))
  
    for _,v := range s {
        m[v]++
    }

    for _,v := range t {
        z[v]++
    }
    
    

    for i := range m {
        if m[i] != z[i] {
            return false
        }
    }

    return true
}
