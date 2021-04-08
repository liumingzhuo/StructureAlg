/*
222. 完全二叉树的节点个数
满二叉树：Perfect Binary Tree  子结点有左右子树
完全二叉树：Complete Binary Tree  子结点靠左排列
Full Binary Tree：要么没有子节点，要么有2个子节点
*/
package completetree222

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	l := root
	r := root
	//左右子树的高度
	var lh, rh int
	for l != nil {
		l = l.Left
		lh++
	}
	for r != nil {
		r = r.Right
		rh++
	}
	//左右子树相等，是一颗满二叉树
	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}
	//否则只是一颗普通二叉树
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
