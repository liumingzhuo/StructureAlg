//25. K 个一组翻转链表

package reversenodes25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a := head
	b := head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newNode := reverseN(a, b)
	a.Next = reverseKGroup(b, k)
	return newNode
}

//反转[a,b)
func reverseN(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	nxt := a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
