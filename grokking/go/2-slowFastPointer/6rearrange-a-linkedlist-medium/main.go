// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/problem-challenge-2-rearrange-a-linkedlist-medium

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type Solution struct{}

// reorder reorders the linked list as per the given logic
func (s *Solution) reorder(head *ListNode) *ListNode {
	// Key Insight:
	// 1. Find the middle of the linked list using the slow and fast pointer technique.
	// 2. Reverse the second half of the linked list starting from the middle.
	// 3. Merge the two halves by alternating nodes from the first half and the reversed second half.

	// Find mid of LL
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	sl, f := head, head
	for f != nil && f.Next != nil {
		prev = sl
		sl = sl.Next
		f = f.Next.Next
	}

	// f!=nil means odd number of nodes
	if f != nil {
		prev = sl
		sl = sl.Next
	}
	prev.Next = nil // Split the linked list into two halves

	// Reverse the second half
	secondHalf := reverseLLFromMid(sl)
	firstHalf := head

	// Merge the two halves
	for firstHalf != nil && secondHalf != nil {
		// Store next pointers
		firstHalfNext := firstHalf.Next
		secondHalfNext := secondHalf.Next

		// Relink: first -> second -> firstNext
		firstHalf.Next = secondHalf
		secondHalf.Next = firstHalfNext

		// Move to next nodes
		firstHalf = firstHalfNext
		secondHalf = secondHalfNext
	}

	return head
}

func reverseLLFromMid(m *ListNode) *ListNode {
	// fmt.Println("Reverse LL from Mid")
	var prev *ListNode
	curr := m
	for curr != nil {
		// fmt.Println("Current", curr, "Prev", prev)
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// fmt.Println("Outside", "Current", curr, "Prev", prev)
	return prev
}

func main() {
	s := &Solution{}
	head := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 10}}}}}
	result := s.reorder(head)
	for result != nil { // Output: 2 10 4 8 6
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Println()
	head1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 10, Next: &ListNode{Val: 12}}}}}}
	result1 := s.reorder(head1)
	for result1 != nil { // Output: 2 12 4 10 6 8
		fmt.Printf("%d ", result1.Val)
		result1 = result1.Next
	}
	fmt.Println()
}
