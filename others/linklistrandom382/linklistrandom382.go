package linklistrandom382

import (
	"math/rand"
	"time"
)

//随机概率算法
//第i个元素被选中的概率是1/i
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution{
		head: head,
		r:    r,
	}
}

/** Returns a random node's value. */
func (s *Solution) GetRandom() int {
	i, res := 0, 0
	p := s.head
	for p != nil {
		i++
		if s.r.Intn(i) == 0 {
			res = p.Val
		}
		p = p.Next
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
