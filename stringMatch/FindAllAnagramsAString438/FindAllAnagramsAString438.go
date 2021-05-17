/*
438. 找到字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
*/
package findallanagramsastring438

import "fmt"

func findAnagrams(s string, p string) []int {
	var (
		left, right, isShrink int
		rltIndexs             []int
	)
	needs := make(map[byte]int, len(s))
	windows := make(map[byte]int, len(s))
	for _, c := range p {
		needs[byte(c)]++
	}
	fmt.Println(needs)
	for right < len(s) {
		rightChar := s[right]
		right++
		if _, ok := needs[rightChar]; ok {
			windows[rightChar]++
			if windows[rightChar] == needs[rightChar] {
				isShrink++
			}
		}
		//缩小窗口的时机为 窗口大于等于 p的大小
		//排列问题 长度需要一致
		for right-left >= len(p) {
			//更新最小范围字串  满足条件
			if isShrink == len(needs) {
				rltIndexs = append(rltIndexs, left)
			}
			leftChar := s[left]
			left++
			if _, ok := needs[leftChar]; ok {
				if windows[leftChar] == needs[leftChar] {
					isShrink--
				}
				windows[leftChar]--
			}
		}
	}
	return rltIndexs
}

func main() {
	findAnagrams("cbaebabacd", "abc")
}
