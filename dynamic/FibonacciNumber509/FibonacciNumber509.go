/*
509. Fibonacci Number
*/
package main

//方法1  备忘录
func fib(n int) int {
	if n < 1 {
		return 0
	}
	//创建一个备忘录
	memo := make([]int, n+1)
	return helper(memo, n)
}
func helper(memo []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}

//自底向上 dp
func fib1(n int) int {
	if n < 1 {
		return 0
	}
	dp := make([]int, n+1)
	if n == 1 || n == 2 {
		return 1
	}
	dp[1] = dp [2] = 1 
	for i := 3; i <= n; i++{
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}

//优化自底向上的dp
func fib2(n int) int{
	if n < 1 {
		return 0
	}
	if n == 1|| n ==2{
		return 1
	}
	curr := 1
	prev := 1
	for i := 3; i <= n; i++{
		sum := curr + prev
        prev = curr
        curr = sum
	}
	return curr
}