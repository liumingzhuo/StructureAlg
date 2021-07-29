/**
[1104] 二叉树寻路
*/
package main

import (
	"math"
	"sort"
)

func pathInZigZagTree(label int) []int {
	// 从底层到label 一共有多少层
	// 每一层的范围是 [2^(n-1),2^n-1]
	// label 上一层是 label//2 的对称点
	// label 上一层是 = 2^n-1 + 2^(n-1) - 对称点
	n, s := 1, 1
	for s*2 <= label {
		n++
		s *= 2
	}
	/** time out */
	// for {
	// 	minValue := math.Pow(2, float64(n-1))
	// 	maxValue := math.Pow(2, float64(n)) - 1
	// 	if label > int(minValue) && label < int(maxValue) {
	// 		break
	// 	}
	// 	n++
	// }
	ans := []int{label}
	for n > 1 {
		n--
		levelTotal := math.Pow(2, float64(n-1)) + (math.Pow(2, float64(n)) - 1)
		parentNode := int(levelTotal) - int(math.Floor(float64(label/2)))
		label = parentNode
		ans = append(ans, parentNode)
	}
	sort.Ints(ans)
	return ans
}
