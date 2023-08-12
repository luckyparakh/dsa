// https://leetcode.com/problems/reverse-linked-list/

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
func (ll *LinkedList) reverseList(head *ListNode) *ListNode {
	node := head
	var prev *ListNode
	for node != nil {
		// fmt.Println("1", node)
		nextNode := node.Next
		// fmt.Println("2", nextNode)
		node.Next = prev
		// fmt.Println("3", node.Next)
		prev = node
		// fmt.Println("4", prev)
		node = nextNode
	}
	ll.head = prev
	return prev
}

func main() {
	ll := initLinkList()
	ll.InsertAtStart(2)
	ll.InsertAtStart(3)
	ll.Traverse()
	ll.reverseList(ll.head)
	ll.Traverse()
}
