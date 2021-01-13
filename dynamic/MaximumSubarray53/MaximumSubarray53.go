/*
53. 最大子序和
*/
package MaximumSubarray53

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	//base case
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < nums[i]+dp[i-1] {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
	}
	res := INT_MIN
	for i := 0; i < n; i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
