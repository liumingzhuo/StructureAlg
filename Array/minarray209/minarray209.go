/*
 * [209] 长度最小的子数组
 */
package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	i, sum := 0, 0
	rltL := math.MaxInt32
	for j := 0; j < len(nums); j++ {
		//扩大窗口
		sum += nums[j]
		//缩小窗口
		for sum >= target {
			subL := j - i + 1
			sum -= nums[i]
			i++
			if subL < rltL {
				rltL = subL
			}
		}
	}
	if rltL == math.MaxInt32 {
		return 0
	}
	return rltL

}

func main() {
	got := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	fmt.Println(got)
}
