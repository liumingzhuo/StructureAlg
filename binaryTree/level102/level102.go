/*
102. 二叉树的层序遍历
*/
package level102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret [][]int

func levelOrder(root *TreeNode) [][]int {
	ret = make([][]int, 0)
	help(root, 0)
	return ret
}
func help(p *TreeNode, level int) {
	if p == nil {
		return
	}
	//层数从0开始，如果等于ret长度，则当前层还没有记录
	if level == len(ret) {
		ret = append(ret, []int{})
	}
	ret[level] = append(ret[level], p.Val)
	help(p.Left, level+1)
	help(p.Right, level+1)
}
