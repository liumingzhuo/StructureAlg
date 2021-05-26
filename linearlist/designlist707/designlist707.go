package main

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  0,
		Next: nil,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index > this.size-1 || index < 0 {
		return -1
	}
	cur := this
	//cur移动到index处
	for ; index >= 0; index-- {
		cur = cur.Next
	}
	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{Val: val}
	newNode.Next = this.Next
	this.Next = newNode
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &MyLinkedList{Val: val}
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	cur := this
	//cur 移动到index之前
	for ; index > 0; index-- {
		cur = cur.Next
	}
	newNode := &MyLinkedList{Val: val}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	cur := this
	// cur 移动到index之前
	for ; index > 0; index-- {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.size--
}
