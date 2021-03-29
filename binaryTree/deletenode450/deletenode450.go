//450. 删除二叉搜索树中的节点
package deletenode450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		//删除结点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//将最大左结点或最小右结点换到删除结点
		//找到右结点中最小的
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
