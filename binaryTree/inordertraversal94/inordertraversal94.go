/**
中序遍历
*/
package inordertraversal94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return res
}
