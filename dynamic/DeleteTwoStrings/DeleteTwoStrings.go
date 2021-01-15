/*
583. 两个字符串的删除操作
*/
package DeleteTwoStrings

var memo [][]int

func minDistance(word1 string, word2 string) int {
	//计算最长公共子序列
	m, n := len(word1), len(word2)
	memo = make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}
	lcs := dp(word1, word2, 0, 0)
	return m - lcs + n - lcs
}

func dp(str1, str2 string, i, j int) int {
	if len(str1) == i || len(str2) == j {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if str1[i] == str2[j] {
		memo[i][j] = dp(str1, str2, i+1, j+1) + 1
	} else {
		memo[i][j] = maxValue(
			dp(str1, str2, i+1, j),
			dp(str1, str2, i, j+1),
		)
	}
	return memo[i][j]
}
func maxValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}
