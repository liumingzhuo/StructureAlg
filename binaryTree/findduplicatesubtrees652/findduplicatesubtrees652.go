package findduplicatesubtrees652

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo = make(map[string]int)
var rlt = make([]*TreeNode, 0)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	traverse(root)
	return rlt
}

//辅助函数
func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.Left)
	right := traverse(root.Right)
	subTree := left + "," + right + "," + strconv.Itoa(root.Val)
	if v, ok := memo[subTree]; ok {
		if v == 1 {
			rlt = append(rlt, root)
		}
		memo[subTree] = v + 1
	} else {
		memo[subTree] = 1
	}
	return subTree

}
