/**
 * [24] 两两交换链表中的节点
 */
package swappairs24

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 -> 2 -> 3
func swapPairs(head *ListNode) *ListNode {
	//创建虚结点  指向head
	vnode := &ListNode{Val: 0, Next: head}
	cur := vnode //指向虚结点的指针
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next            //虚结点后第1个结点
		tmp1 := cur.Next.Next.Next //虚结点之后第三个结点

		cur.Next = cur.Next.Next  // v ->2
		cur.Next.Next = tmp       // v->2->1
		cur.Next.Next.Next = tmp1 // v->2->1->3  此时1 应该是cur.next.next

		cur = cur.Next.Next //交换后 移动cur  2 -> 1(cur) ->3->4
	}
	return vnode.Next
}
