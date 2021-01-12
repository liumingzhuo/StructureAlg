/*
72. 编辑距离 dp 优化
*/
package EditDistance72

func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	//初始化一个m*n的二维切片
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//base case
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	//自底向上
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1]+1,
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
				)
			}
		}
	}
	return dp[m][n]
}
