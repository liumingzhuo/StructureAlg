/*
567. 字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
*/
package PermutationString567

func checkInclusion(s1 string, s2 string) bool {
	left := 0
	right := 0
	needs := make(map[byte]int)
	windows := make(map[byte]int)
	isShrink := 0
	for _, c := range s1 {
		needs[byte(c)]++
	}
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if needs[c] == windows[c] {
				isShrink++
			}
		}
		for right-left >= len(s1) {
			if isShrink == len(needs) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := needs[d]; ok {
				if needs[d] == windows[d] {
					isShrink--
				}
				windows[d]--
			}
		}
	}
	return false
}
