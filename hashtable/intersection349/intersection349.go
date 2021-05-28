/*
 *
 * [349] 两个数组的交集
 */
package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	l := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return l
	}
	//去重后存入m
	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		//如果当前的值 存在切为true 则存入list
		if t, ok := m[v]; ok && t {
			l = append(l, v)
			m[v] = false
		}
	}
	return l
}
func main() {
	got := intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(got)
}
