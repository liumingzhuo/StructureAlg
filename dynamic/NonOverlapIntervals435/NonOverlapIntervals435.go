/*
435. 无重叠区间
*/
package NonOverlapIntervals435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	//边界值
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	e := intervals[0][1] //获取end
	//查找区间
	for _, inter := range intervals {
		s := inter[0]
		if s >= e {
			count++
			e = inter[1]
		}
	}
	return len(intervals) - count
}
