// https://leetcode.com/problems/add-two-numbers/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) InsertAtEnd(v int) {
	n := ListNode{
		Val: v,
	}
	if l.head == nil {
		n.Next = nil
		l.head = &n
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &n
	n.Next = nil
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := initLinkList()
	d := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + d
		r := sum % 10
		ll.InsertAtEnd(r)
		d = sum / 10
		l1 = l1.Next
		l2 = l2.Next

	}
	if l1 == nil {
		for l2 != nil {
			sum := l2.Val + d
			r := sum % 10
			ll.InsertAtEnd(r)
			d = sum / 10
			l2 = l2.Next
		}
	}
	if l2 == nil {
		for l1 != nil {
			sum := l1.Val + d
			r := sum % 10
			ll.InsertAtEnd(r)
			d = sum / 10
			l1 = l1.Next
		}
	}
	if d > 0 {
		ll.InsertAtEnd(d)
	}
	return ll.head
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

func main() {

}
