package maxbinarytree654

//654. 最大二叉树

//TreeNode 结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const intMax = int(^uint(0) >> 1)
const intMin = ^intMax

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, lo, hi int) *TreeNode {
	//base case
	if lo > hi {
		return nil
	}
	//找到数组中的最大值
	index := -1
	maxVal := intMin
	for i := lo; i <= hi; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}
	root := &TreeNode{Val: maxVal}
	root.Left = build(nums, lo, index-1)
	root.Right = build(nums, index+1, hi)
	return root
}
