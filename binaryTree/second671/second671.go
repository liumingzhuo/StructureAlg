/**
[671] 二叉树中第二小的节点
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val > ans && ans != -1 {
			return
		}
		if node.Val > root.Val {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
