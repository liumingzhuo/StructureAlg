/*
 *
 * [429] N 叉树的层序遍历
 */

package nlevel429

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		inner := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			childs := node.Children
			for j := 0; j < len(childs); j++ {
				queue.PushBack(childs[j])
			}
			inner = append(inner, node.Val)
		}
		res = append(res, inner)
	}
	return res

}
