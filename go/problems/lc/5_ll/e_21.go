// https://leetcode.com/problems/merge-two-sorted-lists/
// https://www.youtube.com/watch?v=Xb4slcp1U38&list=PLgUwDviBIf0r47RKH7fdWN54AbWFgGuii&index=4

package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LL struct {
	head *Node
}

func (l *LL) InsertAtStart(v int) {
	n := Node{
		val:  v,
		next: l.head,
	}
	l.head = &n
}

func (l *LL) InsertAtEnd(v int) {
	n := Node{
		val: v,
	}
	if l.head == nil {
		n.next = nil
		l.head = &n
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &n
	n.next = nil
}

func (l *LL) Traverse() {
	node := l.head
	for node != nil {
		fmt.Print(node.val, " ")
		node = node.next
	}
	fmt.Println("")
}

func initLinkList() *LL {
	return &LL{}
}

func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	newLL := initLinkList()
	for list1 != nil && list2 != nil {
		if list1.val < list2.val {
			newLL.InsertAtEnd(list1.val)
			list1 = list1.next
		} else {
			newLL.InsertAtEnd(list2.val)
			list2 = list2.next
		}
	}
	if list1 == nil {
		for list2 != nil {
			newLL.InsertAtEnd(list2.val)
			list2 = list2.next
		}
	}
	if list2 == nil {
		for list1 != nil {
			newLL.InsertAtEnd(list1.val)
			list1 = list1.next
		}
	}
	return newLL.head
}

// Intuition
// list1 should point to smaller number while list2 should point to bigger one
func mergeTwoListsInPlace(list1 *Node, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.val > list2.val {
		list1, list2 = list2, list1
	}

	head := list1
	var tmp *Node
	for list1 != nil && list2 != nil {
		if list1.val < list2.val {
			tmp = list1
			list1 = list1.next
		} else {
			tmp.next = list2
			list1, list2 = list2, list1
		}
	}
	tmp.next = list2
	return head
}

func main() {
	ll := initLinkList()
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(5)
	ll.Traverse()
	ll1 := initLinkList()
	ll1.InsertAtEnd(1)
	ll1.InsertAtEnd(2)
	ll1.Traverse()
	n := mergeTwoListsInPlace(ll.head, ll1.head)
	for n != nil {
		fmt.Print(" ", n)
		n = n.next
	}
	fmt.Println()
}
