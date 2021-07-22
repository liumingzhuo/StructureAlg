package copyrandomlist138

type Node struct {
	Val    int
	Next   *Node
	Random *Node //random指向链表的索引
}

func copyRandomList(head *Node) *Node {
	//创建一个map key为原链表node  value为新链表node
	m := map[*Node]*Node{}
	//cur 原结点
	cur := head
	for cur != nil {
		//创建新的结点
		n := &Node{Val: cur.Val}
		m[cur] = n
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		n := m[cur] //取出带值的node
		n.Next = m[cur.Next]
		n.Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}

/**
方法二：回溯
var m map[*Node]*Node

func back(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := m[node]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	m[node] = newNode
	newNode.Next = back(node.Next)
	newNode.Random = back(node.Random)
	return newNode
}
func copyRandomList(head *Node) *Node {
	m = map[*Node]*Node{}
	return back(head)
}


*/
