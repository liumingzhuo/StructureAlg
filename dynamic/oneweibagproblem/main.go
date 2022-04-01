package main

import "fmt"

func oneweibagproblem(weight []int, value []int, bagweight int) int {
	// 背包重量
	dp := make([]int, bagweight+1)

	for i := 0; i < len(weight); i++ {
		for j := bagweight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	return dp[bagweight]

}

func max(a, b int) int {
	if a > b {
		return a
	}

    return b
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagweight := 4
	maxValue := oneweibagproblem(weight, value, bagweight)
	fmt.Println(maxValue)
}
