//538. 把二叉搜索树转换为累加树
//利用BST中序遍历有序的特性
package bst2greatertree538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	//逆序
	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}
