//3数之和

package main

import (
	"fmt"
	"sort"
)

//先sort 然后使用双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n; i++ {
		//排序后如果第一个元素大于0  则一定不能构成三元组
		if nums[i] > 0 {
			return res
		}
		ts := twoSum(&nums, i+1, -nums[i])
		for _, t := range ts {
			t = append(t, nums[i])
			res = append(res, t)
		}
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
func twoSum(nums *[]int, start, target int) [][]int {
	lo, hi := start, len(*nums)-1
	rlt := [][]int{}
	for lo < hi {
		left, right := (*nums)[lo], (*nums)[hi]
		sum := left + right
		if sum < target {
			for lo < hi && (*nums)[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && (*nums)[hi] == right {
				hi--
			}
		} else {
			rlt = append(rlt, []int{left, right})
			for lo < hi && (*nums)[lo] == left {
				lo++
			}
			for lo < hi && (*nums)[hi] == right {
				hi--
			}
		}
	}
	return rlt
}
func main() {
	rlt := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(rlt)
}
