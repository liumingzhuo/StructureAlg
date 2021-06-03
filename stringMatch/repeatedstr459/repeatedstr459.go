/*
 *
 * [459] 重复的子字符串
 * kmp 在一个串中查找是否出现过另一个串
 */
package main

import "fmt"

// 前缀表减1实现
// func repeatedSubstringPattern(s string) bool {
// 	n := len(s)
// 	if n == 0 {
// 		return false
// 	}
// 	next := make([]int, n)
// 	j := -1
// 	next[0] = j
// 	for i := 1; i < n; i++ {
// 		for j >= 0 && s[i] != s[j+1] {
// 			j = next[j]
// 		}
// 		if s[i] == s[j+1] {
// 			j++
// 		}
// 		next[i] = j
// 	}
// 	// next[n-1]+1 最长相同前后缀的长度
// 	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
// 		return true
// 	}
// 	return false
// }

// 前缀表（不减一）的代码实现
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

func main() {
	got := repeatedSubstringPattern("abac")
	fmt.Println(got)
}
