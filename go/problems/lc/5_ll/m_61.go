// https://leetcode.com/problems/rotate-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// edge case
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Find len of LL
	c := 0
	node := head
	var prev *ListNode
	for node != nil {
		c++
		prev = node
		node = node.Next
	}

	// As number of rotation can be greater than len of ll, hence take modular of k
	k = k % c
	if k == 0 {
		return head
	}

	// Point last node to head
	prev.Next = head

	// Find last Node after rotation
	lastNodePostRotation := head
	lastNodePostRotationIndex := c - k

	for i := 1; i < lastNodePostRotationIndex; i++ {
		lastNodePostRotation = lastNodePostRotation.Next
	}

	// Point last node's next to head
	head = lastNodePostRotation.Next
	lastNodePostRotation.Next = nil
	return head
}

func main() {

}
