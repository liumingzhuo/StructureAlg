package traprainwater42

//trap 接雨水

func trap(height []int) int {

	/*
		//暴力求解
		n := len(height)
		res := 0
		for i := 1; i < n-1; i++ {
			leftMax, rightMax := 0, 0
			//i左边的
			for j := i; j >= 0; j-- {
				if leftMax < height[j] {
					leftMax = height[j]
				}
			}
			//i右侧
			for j := i; j < n; j++ {
				if rightMax < height[j] {
					rightMax = height[j]
				}
			}
			if leftMax <= rightMax {
				res += leftMax - height[i]
			} else {
				res += rightMax - height[i]
			}
		}
		return res
	*/

	//备忘录优化
	res := 0
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i := 0; i < n; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
