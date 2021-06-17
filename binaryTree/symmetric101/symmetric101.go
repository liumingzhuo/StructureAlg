package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
一、使用递归的方法

//是否是对称树，即将一棵树分成左右两棵树，判断左右两棵树是否相等
//一颗树的遍历顺序是左右中， 一棵树的遍历顺序是右左中
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return compare(root.Left, root.Right)
}

//比较函数
func compare(leftNode, rightNode *TreeNode) bool {
	//左树为空，右树为空，则相等
	if leftNode == nil && rightNode == nil {
		return true
	}
	//左树为空，右树不为空，则不相等
	if leftNode == nil && rightNode != nil {
		return false
	}
	//左树不为空，右树为空，则不相等
	if leftNode != nil && rightNode == nil {
		return false
	}
	//左右树的值不相等
	if leftNode.Val != rightNode.Val {
		return false
	}
	//左右树相等， 进入递归判断两棵树的内外层
	outer := compare(leftNode.Left, rightNode.Right)
	inner := compare(leftNode.Right, rightNode.Left)
	return outer && inner

}
*/
/**
二、使用栈
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		//从队列中取出
		leftNode := queue.Remove(queue.Front()).(*TreeNode)
		rightNode := queue.Remove(queue.Front()).(*TreeNode)
		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil || (leftNode.Val != rightNode.Val) {
			return false
		}
		queue.PushBack(leftNode.Left)
		queue.PushBack(rightNode.Right)
		queue.PushBack(rightNode.Left)
		queue.PushBack(leftNode.Right)
	}
	return true
}
