/*
 *
 * [637] 二叉树的层平均值
 */
package averagelevel637

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := list.New()
	queue.PushBack(root)
	//记录每一层的数据
	res := [][]int{}
	for queue.Len() > 0 {
		length := queue.Len()
		inner := []int{}
		for i := 0; i < length; i++ {
			//获取并移除队列首位
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
	fRlt := []float64{}
	for _, val := range res {
		var c float64
		for _, v := range val {
			c += float64(v)
		}
		fRlt = append(fRlt, c/float64(len(val)))
	}
	return fRlt
}
