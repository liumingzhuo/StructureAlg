package diffways2compute241

import "strconv"

//241. 为运算表达式设计优先级
var memo = make(map[string][]int)

func diffWaysToCompute(input string) []int {
	if _, ok := memo[input]; ok {
		return memo[input]
	}
	res := make([]int, 0)
	inputByte := []byte(input)
	for i := 0; i < len(inputByte); i++ {
		c := inputByte[i]
		if c == '+' || c == '-' || c == '*' {
			left := diffWaysToCompute(string(inputByte[:i]))
			right := diffWaysToCompute(string(inputByte[i+1:]))

			//通过字问题的结果 合并
			for _, l := range left {
				for _, r := range right {
					if c == '+' {
						res = append(res, l+r)
					} else if c == '-' {
						res = append(res, l-r)
					} else if c == '*' {
						res = append(res, l*r)
					}
				}
			}
		}
	}
	//递归结束
	if len(res) == 0 {
		num, err := strconv.Atoi("123")
		if err == nil {
			res = append(res, num)
		}
	}
	memo[input] = res
	return res
}
