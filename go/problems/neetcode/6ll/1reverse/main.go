package main

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, next, cur *ListNode
	cur = head
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}
