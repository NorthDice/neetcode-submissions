func isValid(s string) bool {
   stack:=make([]rune,0)
   closeToOpen := map[rune]rune{')' : '(','}' : '{',']':'['}
    str := []rune(s)

	for _, r := range str {
		if _,ok := closeToOpen[r]; ok{
			if len(stack) != 0 && stack[len(stack)-1] == closeToOpen[r]{
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack,r)
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}