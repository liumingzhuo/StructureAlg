/*
1143. 最长公共子序列
*/
package main

var memo [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo = make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}
	return dp(text1, 0, text2, 0)
}
func dp(text1 string, m int, text2 string, n int) int {
	if len(text1) == m || len(text2) == n {
		return 0
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	if text1[m] == text2[n] {
		memo[m][n] = 1 + dp(text1, m+1, text2, n+1)
	} else {
		memo[m][n] = max(
			dp(text1, m+1, text2, n),
			dp(text1, m, text2, n+1),
		)
	}
	return memo[m][n]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	longestCommonSubsequence("abc", "ab")
}
