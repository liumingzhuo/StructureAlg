//98. 验证二叉搜索树
package verifyvalidbst98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isBST(root, nil, nil)
}
func isBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isBST(root.Left, min, root) && isBST(root.Right, root, max)
}
