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
			dp(word1, word2, i, j-1)+1,   //插入
			dp(word1, word2, i-1, j-1)+1, //替换
			dp(word1, word2, i-1, j-1)+1, //替换
		)
	}

}
