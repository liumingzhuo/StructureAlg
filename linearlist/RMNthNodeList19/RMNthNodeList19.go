/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*/
package RMNthNodeList19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//使用快慢指针删除
	fast := head
	slow := head
	//让快指针多走n步
	for ; n > 0; n-- {
		fast = fast.Next
	}
	if fast == nil {
		return slow.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
