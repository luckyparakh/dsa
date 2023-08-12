// https://leetcode.com/problems/intersection-of-two-linked-lists/description/
// Find len of both ll and take diff of len
// Make dummy node of large ll forward by the diff
// After this dummy node would be parallel to start of smaller ll
// Move both by 1 step at a time and keep comparing the nodes

package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA := headA
	lenA := 0
	nodeB := headB
	lenB := 0
	for nodeA != nil {
		lenA++
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		lenB++
		nodeB = nodeB.Next
	}

	nodeA = headA
	nodeB = headB
	if lenA > lenB {
		diff := lenA - lenB

		for i := 0; i < diff; i++ {
			nodeA = nodeA.Next
		}
	} else {
		diff := lenB - lenA

		for i := 0; i < diff; i++ {
			nodeB = nodeB.Next
		}
	}
	for nodeA != nil && nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return &ListNode{}
}
