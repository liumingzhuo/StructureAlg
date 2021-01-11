/*
322. 零钱兑换
*/
package CoinChange322

//定义int 最大值
const INT_MAX = int(^uint(0) >> 1)

//const INT_MIN = ^INT_MAX
func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	return dp(coins, memo, amount)
}

func dp(coins []int, memo map[int]int, n int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n == 0 {
		return 0
	}
	if n == -1 {
		return -1
	}
	res := INT_MAX
	for i := 0; i < len(coins); i++ {
		sub := dp(coins, memo, n-coins[i])
		if sub == -1 {
			continue
		}
		if res > 1+sub {
			res = 1 + sub
		}
	}
	if res != INT_MAX {
		memo[n] = res
	} else {
		memo[n] = -1
	}
	return memo[n]
}
