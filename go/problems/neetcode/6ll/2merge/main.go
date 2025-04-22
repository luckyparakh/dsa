package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type LL struct {
	Head *ListNode
	Cur  *ListNode
}

func (l *LL) insertAtEnd(val int) {
	n := ListNode{
		Val: val,
	}
	if l.Head == nil {
		// fmt.Println("head is nil")
		l.Head = &n
		l.Cur = l.Head
	} else {
		// fmt.Println("head is not nil")
		l.Cur.Next = &n
		l.Cur = &n
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	fmt.Println("started")
	nll := LL{}
	for list1 != nil && list2 != nil {
		fmt.Println("in loop", list1.Val, list2.Val)
		if list1.Val <= list2.Val {
			fmt.Println("less")
			nll.insertAtEnd(list1.Val)
			list1 = list1.Next
			fmt.Println("less end")
		} else {
			fmt.Println("more")
			nll.insertAtEnd(list2.Val)
			list2 = list2.Next
		}
	}
	if list1 != nil {
		for list1 != nil {
			nll.insertAtEnd(list1.Val)
			list1 = list1.Next
		}
	}
	if list2 != nil {
		for list2 != nil {
			nll.insertAtEnd(list2.Val)
			list2 = list2.Next
		}
	}
	// fmt.Println("end")
	return nll.Head
}

func main() {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	mergeTwoLists(&n1, &n2)
}
