/*
 *
 * [454] 四数相加 II
 */
package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	//使用map 存储
	m := make(map[int]int)
	//统计 n1 + n2 和出现的次数
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if v, ok := m[-(n3 + n4)]; ok {
				count += v
			}
		}
	}
	return count
}
func main() {
	n1 := []int{1, 2}
	n2 := []int{-2, -1}
	n3 := []int{-1, 2}
	n4 := []int{0, 2}
	got := fourSumCount(n1, n2, n3, n4)
	fmt.Println(got)
}
