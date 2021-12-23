package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {
	m, n := len(haystack), len(needle)
	// 构造next j是前缀 i是后缀（相同前后缀）
	next := make([]int, n)
	j := -1
	next[0] = j
	// 构造next
	for i := 1; i < n; i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = next[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		next[i] = j
	}
	// 匹配
	j = -1
	for i := 0; i-j <= m; i++ {
		for j >= 0 && haystack[i%m] != needle[j+1] {
			j = next[j]
		}
		if haystack[i%m] == needle[j+1] {
			j++
		}
		if j == n-1 {
			return i - n + 1
		}
	}
	return -1
}

func repeatedStringMatch(a string, b string) int {
	// kmp
	la, lb := len(a), len(b)
	idx := strStr(a, b)
	if idx == -1 {
		return -1
	}
	if la-idx >= lb {
		return 1
	}
	return (lb+idx-la-1)/la + 2
}

func main() {
	l := repeatedStringMatch("abcd", "cdabcdab")
	fmt.Printf("rlt is :%d", l)
}
