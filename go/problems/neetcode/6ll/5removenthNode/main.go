package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fp, sp := head, head
	for range n {
		fp = fp.Next
	}
	if fp == nil {
		head = sp.Next
		sp.Next = nil
		return head
	}
	for fp.Next != nil {
		fp = fp.Next
		sp = sp.Next
	}
	nn := sp.Next
	sp.Next = nn.Next
	nn.Next = nil
	return head
}
