package constbinarytreepos106

//106. 从中序与后序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(inorder []int, inStart, inEnd int,
	postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootValue := postorder[postEnd]
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootValue {
			index = i
			break
		}
	}
	leftSize := index - inStart
	root := TreeNode{Val: rootValue}

	root.Left = build(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1)
	root.Right = build(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return &root
}
