/*
316. 去除重复字母
给你一个字符串 s ，
请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
*/
package main

import "fmt"

func removeDuplicateLetters(s string) string {
	var stack []byte
	inStack := [26]bool{}
	count := [26]int{}
	//统计每个字符的个数
	for _, ch := range s {
		count[ch-'a']++
	}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']--
		if inStack[s[i]-'a'] {
			continue
		}
		//如果当前字符小于前面的字符  则弹出
		for len(stack) > 0 && s[i] < stack[len(stack)-1] {
			//如果要弹出的字符不会出现  则不弹出
			if count[stack[len(stack)-1]-'a'] == 0 {
				break
			}
			//stack[len(stack)-1] 取出最后一个元素
			inStack[stack[len(stack)-1]-'a'] = false
			//等于做pop操作
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		inStack[s[i]-'a'] = true
	}
	return string(stack)
}
func main() {
	rlt := removeDuplicateLetters("bcabc")
	fmt.Println(rlt)
}
