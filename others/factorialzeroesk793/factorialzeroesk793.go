package main

import "fmt"

/*
Example 1:
Input: K = 0
Output: 5
Explanation: 0!, 1!, 2!, 3!, and 4! end with K = 0 zeroes.

Example 2:
Input: K = 5
Output: 0
Explanation: There is no x such that x! ends in K = 5 zeroes.
*/

const uintMax uint32 = ^uint32(0)

func preimageSizeFZF(K int) int {
	return rightBound(uint32(K)) - leftBound(uint32(K)) + 1
}
func leftBound(target uint32) int {
	//在左侧区间中 小于target的个数
	lo, hi := uint32(0), uintMax
	for lo < hi { //左闭右开区间
		mid := lo + (hi-lo)/2
		if trailingZeroes(mid) < target {
			lo = mid + 1 //[mid+1, hi)
		} else if trailingZeroes(mid) > target {
			hi = mid //[lo, mid)
		} else {
			hi = mid //找到后 不要停 继续缩小左侧区间
		}
	}
	return int(lo)
}
func rightBound(target uint32) int {
	lo, hi := uint32(0), uintMax
	for lo < hi {
		mid := lo + (hi-lo)/2
		if trailingZeroes(mid) < target {
			lo = mid + 1
		} else if trailingZeroes(mid) > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return int(hi - 1)
}
func trailingZeroes(n uint32) uint32 {
	var count uint32 = 0
	for i := n; i/5 > 0; i /= 5 {
		count += i / 5
	}
	return count
}

func main() {
	n := trailingZeroes(uintMax) //大于10^9
	fmt.Println(n)
}
