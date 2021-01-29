package jumpgame45

/*
45. 跳跃游戏 II
*/
func jump(nums []int) int {
	n := len(nums)
	var end, farthest, jump int // end 表示跳跃到的位置  farthest表示跳跃最远的位置
	for i := 0; i < n-1; i++ {
		farthest = max(nums[i]+i, farthest) //找到nums[i]这个坐标点，跳最远的位置
		if end == i {                       //跳跃到指定位置
			jump++
			end = farthest
		}
	}
	return jump
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
