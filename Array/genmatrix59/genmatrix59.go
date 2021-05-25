/*
*[59] 螺旋矩阵 II
 */
package main

import "fmt"

func generateMatrix(n int) [][]int {
	startX, startY := 0, 0
	offset := 1
	count := 1
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n/2; i++ {
		x := startX
		y := startY
		//上
		for ; y < startY+n-offset; y++ {
			ret[startX][y] = count
			count++
		}
		//右
		for ; x < startX+n-offset; x++ {
			ret[x][y] = count
			count++
		}
		//下
		for ; y > startY; y-- {
			ret[x][y] = count
			count++
		}
		//左
		for ; x > startX; x-- {
			ret[x][y] = count
			count++
		}
		startX++
		startY++
		offset += 2
	}
	if n%2 != 0 {
		ret[n/2][n/2] = count
	}
	fmt.Println(ret)
	return ret
}
func main() {
	generateMatrix(4)
}
