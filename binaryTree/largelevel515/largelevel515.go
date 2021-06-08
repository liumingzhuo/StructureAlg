package largelevel515

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		inner := []int{}
		for i := 0; i < length; i++ {
			// 获取队列首位node
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			inner = append(inner, node.Val)
		}
		res = append(res, inner)
	}
	//计算每一层中的最大值
	fRlt := []int{}
	for _, data := range res {
		max := math.MinInt64
		for _, d := range data {
			if d > max {
				max = d
			}
		}
		fRlt = append(fRlt, max)
	}
	return fRlt
}
