package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right + 1
}

func main() {
	got := searchInsert([]int{1, 3, 5, 6}, 0)
	fmt.Println(got)
}
