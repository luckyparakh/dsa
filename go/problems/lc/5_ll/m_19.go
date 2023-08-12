// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

import "fmt"

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
func initLinkList() *LinkedList {
	return &LinkedList{}
}
func (l *LinkedList) removeNthFromEnd(head *ListNode, n int) *ListNode {

	c := 1
	var prev *ListNode
	fp, sp := head, head
	for c != n {
		fp = fp.Next
		c++
	}
	// This means node at start to be removed; Edge case
	if fp.Next == nil {
		l.head = sp.Next
		sp.Next = nil
		return l.head
	}

	for fp.Next != nil {
		prev = sp
		sp = sp.Next
		fp = fp.Next
	}
	prev.Next = sp.Next
	sp.Next = nil
	return head
}
func main() {
	ll := initLinkList()
	ll.InsertAtStart(1)
	ll.InsertAtStart(2)
	ll.InsertAtStart(3)
	ll.InsertAtStart(4)
	ll.InsertAtStart(5)
	ll.Traverse()
	ll.removeNthFromEnd(ll.head, 2)
	ll.Traverse()

	ll1 := initLinkList()
	ll1.InsertAtStart(1)
	ll1.InsertAtStart(2)
	ll1.Traverse()
	ll1.removeNthFromEnd(ll1.head, 2)
	ll1.Traverse()

	ll2 := initLinkList()
	ll2.InsertAtStart(1)
	ll2.Traverse()
	ll2.removeNthFromEnd(ll2.head, 1)
	ll2.Traverse()
}
