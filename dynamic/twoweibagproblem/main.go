package main

import "fmt"

// 01背包 二维数组方法
func twoweibagproble(weight []int, value []int, bagweight int) int {
	//dp[i][j] 任选[0, i]上的物品， 放进容量为j的背包中，最大的价值
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	// 初始化dp数组
	// 第0个物品的重量大于背包j
	for j := 0; j < weight[0]; j++ {
		dp[0][j] = 0
	}
	// 初始化物品0
	for j := weight[0]; j <= bagweight; j++ {
		dp[0][j] = value[0]
	}

	// 遍历顺序
	for i := 1; i < len(weight); i++ { // 先遍历物品
		for j := weight[i]; j <= bagweight; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])

		}
	}

	return dp[len(weight)-1][bagweight]
}
func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagweight := 4
	maxValue := twoweibagproble(weight, value, bagweight)
	fmt.Println(maxValue)
}
