/*
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

*/
package mergesorted88

func Merge(nums1 []int, m int, nums2 []int, n int) {
	tail := m + n - 1
	i1 := m - 1
	i2 := n - 1
	for i1 > -1 && i2 > -1 {
		if nums1[i1] < nums2[i2] {
			nums1[tail] = nums2[i2]
			i2--
		} else {
			nums1[tail] = nums1[i1]
			i1--
		}
		tail--
	}
	for i1 > -1 {
		nums1[tail] = nums1[i1]
		i1--
		tail--
	}
	for i2 > -1 {
		nums1[tail] = nums2[i2]
		i2--
		tail--
	}
}
