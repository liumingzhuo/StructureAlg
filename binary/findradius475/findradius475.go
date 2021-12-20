package main

import (
	"fmt"
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0
	for _, house := range houses {
		minDis := math.MaxInt32
		// 第一个大于等于house的heater 坐标
		j := sort.SearchInts(heaters, house+1)
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return ans
}

func main() {
	houses := []int{1, 2, 3}
	heaters := []int{2}
	rlt := findRadius(houses, heaters)
	fmt.Println(rlt)
}
