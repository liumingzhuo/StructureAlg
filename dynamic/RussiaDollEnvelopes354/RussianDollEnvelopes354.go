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
	piles := 0
	for i := 0; i < len(envelopes); i++ {
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] >= envelopes[i][1] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == piles {
			piles++
		}
		//放在堆顶
		top[left] = envelopes[i][1]
	}
	return piles
}
func main() {
	lis := maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}})
	fmt.Println(lis)
}
