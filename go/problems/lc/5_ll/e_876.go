package main

import "fmt"

// https://leetcode.com/problems/middle-of-the-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func (ll *LinkedList) InsertAtStart(val int) {
	node := ListNode{
		Val:  val,
		Next: ll.head,
	}
	ll.head = &node
}
func (ll *LinkedList) Traverse() {
	node := ll.head
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}
func middleNode(head *ListNode) *ListNode {
	sp, fp := head, head
	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}
	return sp
}

func initLinkList() *LinkedList {
	return &LinkedList{}
}
func main() {
	ll := initLinkList()
	ll.InsertAtStart(1)
	ll.InsertAtStart(2)
	ll.InsertAtStart(3)
	ll.InsertAtStart(4)
	ll.InsertAtStart(5)
	ll.Traverse()
	fmt.Println(middleNode(ll.head))
}
