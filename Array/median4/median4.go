/*寻找两个正序数组的中位数
中位数：将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素。
*/
package main

import (
	"fmt"
	"math"
)

//奇数 i + j = m - i + n - j + 1
//偶数 i+j=m−i+n−j
//总有i=(left + right )/2  j = (m+n+1)/2 - i

// A[0]...A[i-1] |  A[i]...A[m-1]
// B[0]...B[j-1] |  B[j]...B[n-1]
//恒有 A[i-1] <= B[j]  ,B[j-1] <= A[i]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//len(nums1) < len(nums2)
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	//
	left, right := 0, m
	m1, m2 := 0, 0
	for left <= right {
		//左区间
		i := (left + right) / 2
		//j 一定是正数
		j := (m+n+1)/2 - i
		num_i_1 := math.MinInt32 //A i-1
		if i != 0 {
			num_i_1 = nums1[i-1]
		}
		num_i := math.MaxInt32
		if i != m {
			num_i = nums1[i]
		}
		num_j_1 := math.MinInt32 //B j-1
		if j != 0 {
			num_j_1 = nums2[j-1]
		}
		num_j := math.MaxInt32
		if j != n {
			num_j = nums2[j]
		}
		//找到最大满足的值
		if num_i_1 <= num_j {
			m1 = max(num_i_1, num_j_1) //前一部分的最大值
			m2 = min(num_i, num_j)     //后一部分的最小值
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)&1 == 0 { //偶数  (max(left_part) + min(right))/2
		return float64(m1+m2) / 2.0
	}
	return float64(m1) //奇数 max(left_part)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	r := findMedianSortedArrays([]int{1, 2}, []int{3})
	fmt.Println(r)
}
