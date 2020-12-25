/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

package LongestCharacters3

func lengthOfLongestSubstring(s string) int {
	windows := make(map[byte]int, len(s))
	var (
		left, right int
		maxLen      int
	)
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}
