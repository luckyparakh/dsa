// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/linkedlist-cycle-easy
package main

// ListNode struct represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution struct is used to group the methods
type Solution struct{}

/*
If the LinkedList has a cycle, the fast pointer enters the cycle first, followed by the slow pointer.
After this, both pointers will keep moving in the cycle infinitely.
If at any stage both of these pointers meet, we can conclude that the LinkedList has a cycle in it.
Let’s analyze if it is possible for the two pointers to meet.
When the fast pointer is approaching the slow pointer from behind we have two possibilities:

	The fast pointer is one step behind the slow pointer.
	The fast pointer is two steps behind the slow pointer.
All other distances between the fast and slow pointers will reduce to one of these two possibilities.

Let’s analyze these scenarios, considering the fast pointer always moves first:
	If the fast pointer is one step behind the slow pointer:
		The fast pointer moves two steps and the slow pointer moves one step, and they both meet.
	If the fast pointer is two steps behind the slow pointer:
		The fast pointer moves two steps and the slow pointer moves one step.
		After the moves, the fast pointer will be one step behind the slow pointer, which reduces this scenario to the first scenario.
		This means that the two pointers will meet in the next iteration.
*/

// hasCycle method checks for a cycle in the linked list
func (s *Solution) hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// findCycleLength finds the length of the cycle in the linked list if it exists.
// If there is no cycle, it returns 0.
// Key Insight:
// Once a cycle is detected (when slow meets fast), we can keep one pointer fixed and move the other pointer
// around the cycle until it meets the fixed pointer again, counting the number of steps taken.
// This count will give us the length of the cycle.
func (s *Solution) findCycleLength(head *ListNode) int {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast { // found the cycle
			return s.calculateLength(slow) // Calculate the cycle length
		}
	}
	return 0
}

// calculateLength calculates the length of the cycle.
func (s *Solution) calculateLength(slow *ListNode) int {
	current := slow
	cycleLength := 0
	for {
		current = current.Next
		cycleLength++
		if current == slow { // Reached the starting point of the cycle
			break
		}
	}
	return cycleLength
}
