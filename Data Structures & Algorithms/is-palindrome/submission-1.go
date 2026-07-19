
func isPalindrome(s string) bool {
	i:=0
	j:=len(s)-1
	s = strings.ToLower(s) 
		for i < j{
		if !isAlphanumeric(s[i]){
			i++
			continue
		}
		if !isAlphanumeric(s[j]){
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return(b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}
