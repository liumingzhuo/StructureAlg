package populatingright116

//Node  结构体结点
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwo(root.Left, root.Right)
	return root
}

//connectTwo 辅助函数 用来连接相邻2个结点
func connectTwo(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2

	connectTwo(node1.Right, node2.Left)
	connectTwo(node1.Left, node1.Right)
	connectTwo(node2.Left, node2.Right)

}
