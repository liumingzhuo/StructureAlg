/*
710. 黑名单中的随机数
给定一个包含 [0，n ) 中独特的整数的黑名单 B，写一个函数从 [ 0，n ) 中返回一个不在 B 中的随机整数。

对它进行优化使其尽量少调用系统方法 Math.random() 。
*/
package RandomBlacklist710

import "math/rand"

type Solution struct {
	sz      int
	numsMap map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	szn := N - len(blacklist)
	mapp := make(map[int]int)
	for _, b := range blacklist {
		mapp[b] = 888
	}
	last := N - 1
	for _, b := range blacklist {
		if b > szn {
			continue
		}
		for _, ok := mapp[last]; ok; {
			last--
		}
		mapp[b] = last
		last--

	}
	return Solution{sz: szn, numsMap: mapp}
}

func (this *Solution) Pick() int {
	i := rand.Intn(this.sz)
	//如果索引命中了黑名单
	if _, ok := this.numsMap[i]; ok {
		return this.numsMap[i]
	}
	return i
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
