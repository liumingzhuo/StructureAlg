/**
使用线段树加扫描线完成 天际线问题
{{x1,height},{x2,0}}
*/
package main

import "fmt"

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	if n == 0 {
		return [][]int{}
	}
	return segment(buildings, 0, n-1)

}

func segment(buildings [][]int, l, r int) (ans [][]int) {
	if l == r {
		return [][]int{{buildings[l][0], buildings[l][2]}, {buildings[l][1], 0}}
	}

	mid := l + (r-l)/2
	left := segment(buildings, l, mid)
	right := segment(buildings, mid+1, r)
	//合并
	m, n := 0, 0
	lh, rh := 0, 0
	for m < len(left) || n < len(right) {
		if m >= len(left) {
			ans = append(ans, right[n])
			n++
		} else if n >= len(right) {
			ans = append(ans, left[m])
			m++
		} else {
			if left[m][0] < right[n][0] {
				//左侧高度大于右侧线
				if left[m][1] > rh {
					ans = append(ans, left[m])
				} else if lh > rh {
					ans = append(ans, []int{left[m][0], rh})
				}
				lh = left[m][1]
				m++
			} else if left[m][0] > right[n][0] {
				if right[n][1] > lh {
					ans = append(ans, right[n])
				} else if rh > lh {
					ans = append(ans, []int{right[n][0], rh})
				}
				rh = right[n][1]
				n++
			} else {
				//左右相同 排除相等的高度
				if left[m][1] >= right[n][1] && left[m][1] != max(lh, rh) {
					ans = append(ans, left[m])
				} else if left[m][1] <= right[n][1] && right[n][1] != max(lh, rh) {
					ans = append(ans, right[n])
				}
				lh = left[m][1]
				rh = right[n][1]
				m++
				n++
			}
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	// want := [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}
	got := getSkyline(buildings)
	fmt.Println(got)
}
