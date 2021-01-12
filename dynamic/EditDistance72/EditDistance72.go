/*
72. 编辑距离
*/
package EditDistance72

const MAX = int(^uint(0) >> 1)

func minDistance(word1 string, word2 string) int {
	return dp(word1, word2, len(word1)-1, len(word2)-1)
}

func min(a ...int) int {
	m := MAX
	for _, i := range a {
		if i < m {
			m = i
		}
	}
	return m
}

//递归解法  dp返回 word1[0...i] 和word[0...j]的最短距离
func dp(word1 string, word2 string, i int, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if word1[i] == word2[j] {
		return dp(word1, word2, i-1, j-1)
	} else {
		return min(
			dp(word1, word2, i, j-1)+1,   //插入 在word1插入一个和word2同样的字符，那么word2[j]就被匹配了  向前走j
			dp(word1, word2, i-1, j-1)+1, //替换
			dp(word1, word2, i-1, j)+1,   //删除 直接把word1[i]这个字符删除
		)
	}

}
