/*
 *
 * [18] 四数之和
 */
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	//排序
	sort.Ints(nums)
	//结果
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				s := nums[i] + nums[j] + nums[left] + nums[right]
				if s < target {
					left++
				} else if s > target {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					//去重
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}
func main() {
	got := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	// got := fourSum([]int{}, 0)
	fmt.Println(got)
}
