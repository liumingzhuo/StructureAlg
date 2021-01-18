/*
354. 俄罗斯套娃信封问题
*/

package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	//先对行进行排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	top := make([]int, len(envelopes))
	start, end, rlt := 0, 0, 0
	for i := 0; i < len(envelopes); i++ {
		start = 0
		end = rlt
		for start < end {
			mid := start + (end-start)/2
			if top[mid] >= envelopes[i][1] {
				end = mid
			} else {
				start = mid + 1
			}
		}
		if start == rlt {
			rlt++
		}
		top[start] = envelopes[i][1]
	}
	return rlt
}
func main() {
	lis := maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}})
	fmt.Println(lis)
}
