/**
Boyer-Moore 投票算法

数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。
请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
*/
package main

import "fmt"

func majorityElement(nums []int) int {
	ele := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			ele = num
		}
		if ele == num {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if ele == num {
			count++
		}
		if count*2 > len(nums) {
			return ele
		}
	}
	return -1
}

func main() {
	input := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	got := majorityElement(input)
	fmt.Println(got)
}
