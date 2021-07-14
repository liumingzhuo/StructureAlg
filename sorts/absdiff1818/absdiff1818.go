/**
对排序好的数据 进行二分查找
*/
package main

import (
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sorted := append(sort.IntSlice(nil), nums1...)
	sorted.Sort()
	n := len(nums1)
	sum := 0
	maxd := 0
	for i := range nums2 {
		a, b := nums1[i], nums2[i]
		if a == b {
			continue
		}
		diff := abs(a - b)
		//二分法找最接近的值
		sum += diff
		l, r := 0, n-1
		for l < r {
			mid := l + (r-l)/2
			if sorted[mid] <= b {
				l = mid + 1
			} else {
				r = mid
			}
		}
		//找到最小的大于等于b的数，但不一定最接近b
		diffn := abs(sorted[l] - b)
		//如果l左边的数 更接近b 那么取左边
		if l-1 > 0 {
			diffn = min(diffn, abs(b-sorted[l-1]))
		}
		if diffn < diff {
			maxd = max(maxd, diff-diffn)
		}
	}
	//结果中 sum - maxd
	return (sum - maxd) % (1e9 + 7)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
