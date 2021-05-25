/*有序数组的平方 O(n)*/
package main

import (
	"fmt"
	"math"
)

//[-4,-1,0,3,10]
//[0,1,9,16,100]
func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	rlt := make([]int, n)
	index := n - 1
	for left <= right {
		if math.Abs(float64(nums[left])) > math.Abs(float64(nums[right])) {
			rlt[index] = nums[left] * nums[left]
			left++
		} else {
			rlt[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return rlt
}
func main() {
	r := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println(r)
}
