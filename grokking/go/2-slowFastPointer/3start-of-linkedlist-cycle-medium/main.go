// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/start-of-linkedlist-cycle-medium

package main

// ListNode struct definition
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution struct definition
type Solution struct{}

// findCycleStart method of Solution struct
func (s *Solution) findCycleStart(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// fmt.Println(fast.Val, slow.Val)
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
// Key Insight:
// Once a cycle is detected (when slow meets fast), we can reset one pointer to the head of the linked list.
// Then, we move both pointers one step at a time.
// The point at which they meet again will be the start of the cycle.
// This works because the distance from the head to the start of the cycle is equal to the distance from the meeting point to the start of the cycle when moving within the cycle.
// This is classic Floyd's Tortoise and Hare algorithm extension to find the start of the cycle.