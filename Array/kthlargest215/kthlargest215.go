package kthlargest215

import (
	"math/rand"
)

//快速选择排序
func findKthLargest(nums []int, k int) int {
	randNums(nums)
	lo := 0
	hi := len(nums) - 1
	k = len(nums) - k
	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func partition(nums []int, lo, hi int) int {
	if lo >= hi {
		return lo
	}
	pivot := nums[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[hi] = nums[hi], nums[i+1]
	return i + 1

	/*
		//本解法超时
		for {
			for nums[i] < pivot {
				if i == hi {
					break
				}
				i++
			}
			for nums[j] > pivot {
				if j == lo {
					break
				}
				j--
			}
			if i >= j {
				break
			}
			//调换i 和 j的位置
			nums[i], nums[j] = nums[j], nums[i]
			//nums[lo..i] < pivot < nums[j..hi]
		}
		nums[j], nums[lo] = nums[lo], nums[j]
		//最后将j和lo 换位置，保证pivot在中间
		//排序前  4 1 6 3 2 5  //pivot 是4
		//排序后  3 1 2 4 6 5
		return j
	*/
}

func randNums(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		r := rand.Intn(n - i)
		nums[i], nums[r] = nums[r], nums[i]
	}
}

// func randNums(nums *[]int) {
// 	n := len(*nums)
// 	for i := 0; i < n; i++ {
// 		r := rand.Intn(n - i)
// 		(*nums)[i], (*nums)[r] = (*nums)[r], (*nums)[i]
// 	}
// }
