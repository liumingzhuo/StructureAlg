package main

import "math/bits"

func maxLength(arr []string) int {

	//存储符合条件 不重复的字符串
	l := make([]int, 0)
outer:
	//使用二进制存储每一个字符
	for _, s := range arr {
		//当前字符串的二进制
		cur := 0
		for _, ch := range s {
			// 字符的位移位数
			p := ch - 'a'
			//字符串的二进制右移p位，
			//如果存在，则说明此字符已经存在于cur中，不符合要求
			if cur>>p&1 == 1 {
				continue outer
			}
			//不存在 则加入到cur中
			cur = cur | 1<<p
		}
		l = append(l, cur)
	}
	ans := 0
	var backtrack func(int, int)
	backtrack = func(pos, cur int) {
		if pos == len(l) {
			ans = max(ans, bits.OnesCount64(uint64(cur)))
			return
		}
		//判断当前cur和l[pos]是否相同，相同则重复
		//不相同 则拼接两个字符串
		if cur&l[pos] == 0 {
			backtrack(pos+1, cur|l[pos])
		}
		backtrack(pos+1, cur)

	}
	backtrack(0, 0)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tests := []string{"un", "iq", "ue"}
}
