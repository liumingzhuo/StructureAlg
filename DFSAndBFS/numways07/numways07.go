/**
n = 5, relation = [[0,2],[2,1],[3,4],[2,3],[1,4],[2,0],[0,4]], k = 3

*/
package main

import (
	"container/list"
	"fmt"
)

/**
//使用dfs
func numWays(n int, relation [][]int, k int) int {
	//定义一个字典，放入站点和对应的下一个站点
	stations := make(map[int][]int)
	for _, r := range relation {
		c, n := r[0], r[1]
		stations[c] = append(stations[c], n)
	}
	var dfs func(int, int)
	var ans int = 0
	dfs = func(step int, cur int) {
		// 到达k步
		if step == k {
			if cur == n-1 {
				ans++
			}
			return
		}
		// 遍历当前点 能到达的所有点
		for _, s := range stations[cur] {
			dfs(step+1, s)
		}
	}
	dfs(0, 0)
	return ans
}
*/

//使用bfs
func numWays(n int, relation [][]int, k int) (ans int) {
	//定义一个字典，放入站点和对应的下一个站点
	stations := make(map[int][]int)
	for _, r := range relation {
		stations[r[0]] = append(stations[r[0]], r[1])
	}
	queue := list.New()
	queue.PushBack(0)
	for queue.Len() > 0 && k > 0 {
		s := queue.Len()
		for i := 0; i < s; i++ {
			first := queue.Remove(queue.Front()).(int)
			if nss, ok := stations[first]; ok {
				for _, ns := range nss {
					queue.PushBack(ns)
				}
			}
		}
		k--
	}
	for queue.Len() > 0 {
		first := queue.Remove(queue.Front()).(int)
		if first == n-1 {
			ans++
		}
	}
	return ans
}
func main() {
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	rlt := numWays(5, relation, 3)
	fmt.Println(rlt)
}
