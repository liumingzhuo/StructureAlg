package rightpoint117

import (
	"container/list"
)

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
	res := [][]*Node{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		inner := []*Node{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			inner = append(inner, node)
		}
		res = append(res, inner)
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i])-1; j++ {
			res[i][j].Next = res[i][j+1]
		}
	}
	return root
}
