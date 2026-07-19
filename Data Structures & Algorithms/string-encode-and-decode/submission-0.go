type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var b strings.Builder

    for _, v := range strs {
        b.WriteString(strconv.Itoa(len(v))) 
        b.WriteByte('#')
        b.WriteString(v)
    }

    return b.String()
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    i := 0 
    for i < len(encoded) {
        j := i 
        for encoded[j] != '#' {
            j++
        }
        
        length, _ := strconv.Atoi(encoded[i:j])

        i = j + 1

        word := encoded[i : i+length]
        res = append(res, word)

        i += length
    }
    return res
}