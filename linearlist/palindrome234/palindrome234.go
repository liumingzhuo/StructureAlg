package palindrome234

type ListNode struct {
	Val  int
	Next *ListNode
}

var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)

}
func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	res = res && (right.Val == left.Val)
	left = left.Next
	return res

}
