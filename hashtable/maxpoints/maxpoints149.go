package main

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	/**
	方法一
	n := len(points)
	if n < 3 {
		return n
	}
	ans := 0
	for i := 0; i < n; i++ {
		// 计算每2个点的斜率
		for j := i + 1; j < n; j++ {
			//当前最长的点数
			count := 2
			x1 := points[i][0] - points[j][0]
			y1 := points[i][1] - points[j][1]
			//继续比较
			for k := j + 1; k < n; k++ {
				//1/x1 == (points[i][1]-points[k][1])/(points[i][0]-points[k][0])
				if y1*(points[i][0]-points[k][0]) == x1*(points[i][1]-points[k][1]) {
					count++
				}
			}
			ans = max(ans, count)
		}
	}
	return ans
	*/

	//简化时间复杂度
	n := len(points)
	if n < 3 {
		return n
	}
	ans := 2
	//存储斜率
	for i := 0; i < n; i++ {
		slops := make(map[float32]int)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			y1 := points[i][1] - points[j][1]
			x1 := points[i][0] - points[j][0]
			var slop float32
			//x1 = 0  -> slop = +Inf
			slop = float32(y1) * 1.0 / float32(x1)
			count, ok := slops[slop]
			if ok {
				count++
				slops[slop] = count
			} else {
				slops[slop] = 2
			}
			ans = max(ans, count)
		}
	}
	return ans

}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	//[[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
	points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	rlt := maxPoints(points)
	fmt.Println(rlt)

}
