/*
*[203] 移除链表元素
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	//设置一个虚拟头结点
	vhead := &ListNode{
		Next: head,
	}
	cur := vhead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return vhead.Next
}
