package randompickindex398

import (
	"math/rand"
	"time"
)

//398. 随机数索引

type Solution struct {
	arr []int
	r   *rand.Rand
}

func Constructor(nums []int) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution{
		arr: nums,
		r:   r,
	}
}

func (s *Solution) Pick(target int) int {
	k, res := 0, 0
	for i := 0; i < len(s.arr); i++ {
		if s.arr[i] == target {
			k++
			if s.r.Intn(k) == 0 {
				res = i
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
