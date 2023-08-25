// https://leetcode.com/problems/reverse-nodes-in-k-group/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	c := 1
	sp := head
	ep := head
	var prev *ListNode
	for ep != nil {
		// Start c from 1 to avoid hitting below condition, very 1st time
		// and also at this moment Window size is not equal to k (except if k==1)
		if c%k == 0 {
			// Save next node address because ep.Next will change after reverse func
			nn := ep.Next
			reverseHead := reverse(sp, ep.Next)
			if sp == head {
				head = reverseHead
			} else {
				prev.Next = reverseHead
			}
			// After reverse sp is the end of that group of k nodes, hence sp next should point to start of next k group of nodes
			prev = sp
			// Both sp, ep should point to nn
			sp = nn
			ep = nn
		} else {
			ep = ep.Next
		}
		c++
	}
	return head
}
func reverse(s, e *ListNode) *ListNode {
	// If we reverse a LL then its start node's next address should point to end
	prev := e
	for s != e {
		nextNode := s.Next
		s.Next = prev
		prev = s
		s = nextNode
	}
	return prev
}
func reverseKGroup1(head *ListNode, k int) *ListNode {
	// Create a dummy whose next is head
	dummy := &ListNode{Next: head}
	groupPrev := dummy
	for {
		// find kth node, if kth node is nil means end of list or remaining nodes are less than k
		kNode := kNode(groupPrev, k)
		if kNode == nil {
			break
		}
		// Prev node of group will 
		prev, curr := kNode.Next, groupPrev.Next
		for curr != kNode.Next {
			nn := curr.Next
			curr.Next = prev
			prev = curr
			curr = nn
		}
		tmp := groupPrev.Next
		groupPrev.Next = kNode
		groupPrev = tmp
	}
	return head
}
func kNode(start *ListNode, k int) *ListNode {
	for i := 0; i < k; i++ {
		if start == nil {
			return nil
		}
		start = start.Next
	}
	return start
}
func main() {

}
