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
	mapp := make(map[int]int, 0)
	bm := make(map[int]struct{})
	notInBlack := make([]int, 0)
	w := N - len(blacklist)
	flag := 0
	//将黑名单中大于分界点的值 加入到bm中
	for _, b := range blacklist {
		if b >= w {
			bm[b] = struct{}{}
		}
	}
	//将分界点右侧不在黑名单中到值  加入到notInBlack
	for i := w; i < N; i++ {
		if _, ok := bm[i]; !ok {
			notInBlack = append(notInBlack, i)
		}
	}
	//将分界点左侧黑名单中的值 映射到notInBlack中的值
	for _, b := range blacklist {
		if b < w {
			mapp[b] = notInBlack[flag]
			flag++
		}
	}
	return Solution{sz: w, numsMap: mapp}
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
