package bracketsvalid20

//有效的括号
func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && leftis(s[i]) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//根据左括号 返回右括号
func leftis(c byte) byte {
	switch c {
	case ')':
		return '('
	case '}':
		return '{'
	default:
		return '['
	}
}
