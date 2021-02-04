package mindepth111

//111. 二叉树的最小深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 1
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			//从queue取出node
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}
	return depth
}
