package levelbottom107

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	rlt := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		inner := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			inner = append(inner, node.Val)
		}
		rlt = append(rlt, inner)
	}
	//反转
	length := len(rlt)
	for i := 0; i < length/2; i++ {
		rlt[i], rlt[length-i-1] = rlt[length-i-1], rlt[i]
	}
	return rlt
}
