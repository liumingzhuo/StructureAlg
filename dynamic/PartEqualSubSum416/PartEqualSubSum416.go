/*
416. 分割等和子集
*/
package main

/*
背包问题
1. 状态：背包的容量和可选择的物品
2. 选择：可装进背包和不可装进背包
*/

func canPartition(nums []int) bool {
	/*
		//分2个背包
		s := 0
		for _, n := range nums {
			s += n
		}
		if s%2 != 0 {
			return false
		}
		s = s / 2
		n := len(nums)
		dp := make([][]bool, n+1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]bool, s+1)
			for j := 0; j <= s; j++ {
				dp[i][j] = false
			}
			dp[i][0] = true
		}

		for i := 1; i <= n; i++ {
			for j := 1; j <= s; j++ {
				if j-nums[i-1] < 0 {
					//放不进去
					dp[i][j] = dp[i-1][j]
				} else {
					//放或者不放
					dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
				}
			}
		}
		return dp[n][s]*/
	n, sum := len(nums), 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	var dp []bool = make([]bool, sum+1)
	//base case
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[sum]
}
func main() {
	canPartition([]int{1, 2, 3, 4})
}
