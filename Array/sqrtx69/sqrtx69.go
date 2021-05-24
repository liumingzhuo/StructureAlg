package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid == x/mid {
			return mid
		}
		if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//返回right
	return right
}

func main() {
	got := mySqrt(2)
	fmt.Println(got)
}
