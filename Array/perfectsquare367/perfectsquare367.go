package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)/2
		
    if mid*mid == num {
			return true
		}
		if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
func main() {
	got := isPerfectSquare(5)
	fmt.Println(got)
}
