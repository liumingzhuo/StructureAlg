package jumpgame

/*
55. 跳跃游戏
*/

//贪心算法
func canJump(nums []int) bool {
	n := len(nums)
	jump := 0
	for i := 0; i < n-1; i++ {
		jump = max(jump, i+nums[i])
		if jump <= i {
			return false
		}
	}
	return jump >= n-1
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
