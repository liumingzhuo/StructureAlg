package bascalculator225

import "container/list"

//224. 基本计算器
func calculate(s string) int {
	stack := list.New()
	res, sign := 0, 1
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			num = 10*num + int(s[i]-'0')
		} else if s[i] == '+' {
			res = res + sign*num
			sign = 1
			num = 0
		} else if s[i] == '-' {
			res = res + sign*num
			sign = -1
			num = 0
		} else if s[i] == '(' {
			stack.PushBack(res)
			stack.PushBack(sign)
			res = 0
			sign = 1
		} else if s[i] == ')' {
			res = res + sign*num
			res = res * stack.Remove(stack.Back()).(int)
			res = res + stack.Remove(stack.Back()).(int)
			num = 0
		}
	}
	return res + num*sign
}
