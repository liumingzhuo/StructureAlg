/*
 *
 * [54] 螺旋矩阵
 */
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return make([]int, 0)
	}
	var (
		m                        = len(matrix)
		n                        = len(matrix[0])
		top, right, bottom, left = 0, n - 1, m - 1, 0
		index                    = 0
		rlt                      = make([]int, m*n)
	)
	for left <= right && top <= bottom {
		//[left, right]
		for i := left; i <= right; i++ {
			rlt[index] = matrix[top][i]
			index++
		}
		//(top,bottom]
		for i := top + 1; i <= bottom; i++ {
			rlt[index] = matrix[i][right]
			index++
		}
		if left < right && top < bottom {
			//(right,left]
			for i := right - 1; i >= left; i-- {
				rlt[index] = matrix[bottom][i]
				index++
			}
			//(bottom, top)
			for i := bottom - 1; i > top; i-- {
				rlt[index] = matrix[i][left]
				index++
			}
		}

		top++
		right--
		bottom--
		left++
	}
	return rlt
}
func main() {
	got := spiralOrder([][]int{
		// {1, 2, 3},
		// {4, 5, 6},
		// {7, 8, 9},
		{1, 2, 3, 5},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})
	fmt.Println(got)
}
