/**
*852. 山脉数组的峰顶索引
 */
package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-2
	// res := 0
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	arr := []int{3, 4, 5, 1}
	a := peakIndexInMountainArray(arr)
	fmt.Println(a)
}
