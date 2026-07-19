func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
    for i := 0; i < len(strs); i++ {
    var count [26]int
    s := strs[i] 
    for j := 0; j < len(s); j++ {
        count[s[j]-'a']++
    }
    m[count] = append(m[count], s)
}
result := [][]string{}
for _,group := range m {
result = append(result, group)
}

return result

}
    


