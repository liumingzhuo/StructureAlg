/**
02.07. 链表相交
*/
package intersection0207

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	alen, blen := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		alen++
		curA = curA.Next
	}
	for curB != nil {
		blen++
		curB = curB.Next
	}
	curA = headA
	curB = headB
	if blen > alen { //保证A 大于B
		curA, curB = curB, curA
		alen, blen = blen, alen
	}
	//移动A 到与B尾部对其
	l := alen - blen
	for l > 0 {
		curA = curA.Next
		l--
	}
	for curA != nil && curA != curB {
		curA = curA.Next
		curB = curB.Next
	}
	return curA

}
