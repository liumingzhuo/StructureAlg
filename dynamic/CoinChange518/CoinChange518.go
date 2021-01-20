/*
518. 零钱兑换 II
*/
package CoinChange518

func change(amount int, coins []int) int {
	/*n := len(coins)
	dp := make([][]int, n+1) // 存储第i个物品第价值j
	for i := range dp {
		dp[i] = make([]int, amount+1) // 记录第i个硬币价值j
	}
	//base case   当需要价值为0  则只有1种方法
	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n][amount]*/

	//压缩状态
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= n; i++ { //从i从1开始
		for j := 1; j <= amount; j++ { //金额
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[amount]
}
