/*
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/
package minwindowsub76

func minWindow(s string, t string) string {
	needs := make(map[byte]int, len(t))
	windows := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	left := 0
	right := 0
	isShrink := 0
	//最小覆盖字串的起始
	startIndex := 0
	//最小覆盖字串的长度
	lenStr := 999999999
	for right < len(s) {
		c := s[right]
		//右移窗口
		right++
		//更新窗口数据
		if _, ok := needs[c]; ok {
			windows[c]++
			if windows[c] == needs[c] {
				isShrink++
			}
		}
		//判断左窗口是否收缩
		for isShrink == len(needs) {
			//更新最小覆盖字串
			if right-left < lenStr {
				startIndex = left
				lenStr = right - left
			}
			d := s[left]
			left++
			//更新窗口数据
			if _, ok := needs[d]; ok {
				if windows[d] == needs[d] {
					isShrink--
				}
				windows[d]--
			}
		}
	}
	if lenStr == 999999999 {
		return ""
	} else {
		return s[startIndex : startIndex+lenStr]
	}
}
