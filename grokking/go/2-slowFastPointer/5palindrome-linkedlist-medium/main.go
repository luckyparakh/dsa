// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/problem-challenge-1-palindrome-linkedlist-medium

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

func (sl *Solution) isPalindrome(head *ListNode) bool {
	// Key Insight:
	// We can use the slow and fast pointer technique to find the middle of the linked list.
	// Once we find the middle, we can reverse the second half of the linked list.
	// Then, we can compare the first half and the reversed second half node by node to check for palindrome.
	// If all corresponding nodes are equal, the linked list is a palindrome.

	// Edge case: empty list or single node list is a palindrome
	if head == nil || head.Next == nil {
		return true
	}
	mid := findMid(head)
	fmt.Println("Mid found at value", mid.Val)
	headReversedLL := reverseLLFromMid(mid)
	fmt.Println("Reversed LL from mid", headReversedLL.Val)
	isPalin := false

	head1 := head
	headR := headReversedLL
	for head1 != nil && headR != nil {
		if head1.Val != headR.Val {
			break
		}
		head1 = head1.Next
		headR = headR.Next
	}

	if (head1 == nil && headR == nil) || (head1 != nil && headR == nil) {
		isPalin = true
	}

	return isPalin
}

func reverseLLFromMid(m *ListNode) *ListNode {
	fmt.Println("Reverse LL from Mid")
	var prev *ListNode
	curr := m
	for curr != nil {
		fmt.Println("Current", curr, "Prev", prev)
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	fmt.Println("Outside", "Current", curr, "Prev", prev)
	return prev
}

func findMid(head *ListNode) *ListNode {
	fmt.Println("Finding Mid")
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	// If f is not nil, the length of LL is odd, so we return s.Next
	// If f is nil, the length of LL is even, so we return s
	// For example
	// 1->2->3->2->1 (odd length) => mid is 3, we return 2 (s.Next)
	// 1->2->2->1 (even length) => mid is 2 (first 2), we return 2 (s)
	if f != nil {
		fmt.Println("Mid found, returning next of mid for odd length LL")
		return s.Next
	}
	fmt.Println("Mid found, returning mid for even length LL")
	return s
}

func main() {
	// Example usage:
	s := Solution{}
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(s.isPalindrome(head)) // true
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	fmt.Println(s.isPalindrome(head1)) // true
	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{1, &ListNode{1, nil}}}}}
	fmt.Println(s.isPalindrome(head2)) // false
	head3 := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(s.isPalindrome(head3)) // false
	head4 := &ListNode{1, &ListNode{2, &ListNode{1, nil}}}
	fmt.Println(s.isPalindrome(head4)) // true
	head5 := &ListNode{1, nil}
	fmt.Println(s.isPalindrome(head5)) // true
}
