package constbinarytreepre105

//105. 从前序与中序遍历序列构造二叉树

//TreeNode 结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]
	//rootVal 在中序遍历中的索引
	index := 0
	for i := inStart; i < inEnd; i++ {
		if rootVal == inorder[i] {
			index = i
			break
		}
	}
	leftSize := index - inStart

	root := &TreeNode{Val: rootVal}
	root.Left = build(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)
	return root

}
