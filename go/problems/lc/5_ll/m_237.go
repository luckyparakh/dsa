// https://leetcode.com/problems/delete-node-in-a-linked-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Copy value of next node in current node

func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}
func main() {

}
