// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/middle-of-the-linkedlist-easy
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

// findMiddle finds the middle node in a linked list
func (ss *Solution) findMiddle(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}
