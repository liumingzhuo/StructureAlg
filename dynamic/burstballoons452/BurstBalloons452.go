package burstballoons452

import "sort"

/*
 452. 用最少数量的箭引爆气球
*/
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	xEnd := points[0][1]
	for _, p := range points {
		start := p[0]
		if start > xEnd { //如果2个边界相等  则为爆炸
			count++
			xEnd = p[1]
		}
	}
	return len(points) - count
}
