/*
双指针
*/
package traprainwater42

func trapp(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	//定义左右指针
	left := 0
	right := n - 1

	res := 0

	//定义指针左侧最大值
	lMax := height[left]
	//定义指针右侧最大值
	rMax := height[right]

	for left <= right {
		lMax = MaxInt(lMax, height[left])
		rMax = MaxInt(rMax, height[right])

		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}
	return res

}
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
