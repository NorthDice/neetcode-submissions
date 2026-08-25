func isValid(s string) bool {
  closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}
  var stack []rune

  for _, r := range s {
    if openBracket, isClosing := closeToOpen[r]; isClosing {
        if len(stack) == 0 || stack[len(stack)-1] != openBracket {
            return false
        }
        stack = stack[:len(stack)-1]
    } else {
        stack = append(stack, r)
    }
  }
  return len(stack) == 0
}