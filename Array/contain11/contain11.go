/*
盛最多水的容器
*/
package main

import "fmt"

//双指针  存水的大小和左右边界高度和 左右间隔有关
func maxArea(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left, right := 0, n - 1
	res := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	areas := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(areas)
}
