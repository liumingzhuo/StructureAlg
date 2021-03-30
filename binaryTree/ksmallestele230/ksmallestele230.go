//230. 二叉搜索树中第K小的元素
package ksmallestele230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var rlt, cur int

func kthSmallest(root *TreeNode, k int) int {
	rlt = 0
	cur = 0
	traverse(root, k)
	return rlt
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	cur++
	if cur == k {
		rlt = root.Val
		return
	}
	traverse(root.Right, k)
}
