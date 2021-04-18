//104. 二叉树的最大深度
package maxdepth104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
