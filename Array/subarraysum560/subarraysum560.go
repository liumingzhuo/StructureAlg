package subarraysum560

//560. 和为K的子数组
//前缀和
func subarraySum(nums []int, k int) int {
	n := len(nums)
	preSum := make(map[int]int, n+1)
	preSum[0] = 1
	res, sum_i := 0, 0
	for i := 0; i < n; i++ {
		//记录前n项的和
		sum_i += nums[i]
		//计算k-sum_i
		sum_j := sum_i - k
		//存在前缀和 则直接更新结果
		if v, ok := preSum[sum_j]; ok {
			res += v
		}
		//更新preSum 并更新出现的次数
		preSum[sum_i] = preSum[sum_i] + 1
	}
	return res
}
