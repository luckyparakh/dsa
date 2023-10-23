// https://leetcode.com/problems/merge-k-sorted-lists/
// https://www.youtube.com/watch?v=RYJ3vs8qx10

// Package main provides a solution to the problem of merging k sorted lists.
package main

import (
	"container/heap"
	"fmt"
)

// ListNode represents a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MinHeap represents a min heap of ListNodes.
type MinHeap []*ListNode

// Len returns the length of the MinHeap.
func (h MinHeap) Len() int {
	return len(h)
}

// Swap swaps the elements at indices i and j.
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less returns true if the element at index i is less than the element at index j.
func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

// Pop removes and returns the element at the top of the MinHeap.
func (h *MinHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

// Push adds an element x to the MinHeap.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

/*
The function  `mergeKLists`  takes in a slice of linked lists  `lists`  and merges them into a single sorted linked list. 
It uses a min heap  `mh`  to keep track of the minimum node from each linked list.
First, it initializes the  `head`  and  `tail`  pointers to  `nil` . It also creates an empty min heap  `mh` .
Then, it iterates through each linked list in  `lists` . 
If a linked list is not  `nil` , it pushes it into the min heap using the  `Push`  method.
Next, it enters a loop that continues until the min heap is empty. 
In each iteration, it pops the smallest node from the min heap using the  `Pop`  method and assigns it to the variable  `node` .
If the popped node has a  `Next`  node, it pushes the  `Next`  node into the min heap.
The  `Next`  pointer of the popped node is set to  `nil`  to disconnect it from the rest of the linked list.
If the  `head`  pointer is  `nil` , it means that the result linked list is currently empty. 
In this case, the  `head`  and  `tail`  pointers are set to the popped node.
If the  `head`  pointer is not  `nil` , it means that the result linked list already has nodes. 
In this case, the popped node is appended to the end of the result linked list by setting the  `Next`  pointer of the  `tail`  node to the popped node, and updating the  `tail`  pointer to the popped node.
Finally, the function returns the  `head`  pointer, which points to the head of the merged linked list.


Therefore, the overall time complexity of the code is O(N log k).
*/
func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	mh := MinHeap{}
	
	// Push non-nil lists into the min heap.
	// TC: O(k log k) as only 1st element of k lists is push in heap, and pushing take log K
	for _, l := range lists {
		if l != nil {
			heap.Push(&mh, l)
		}
	}
	
	// Pop the smallest node from the min heap and append it to the result list.
	for mh.Len() > 0 {
		// TC O( N Log K) all elements will be taken out and pop takes O(log k)
		node := heap.Pop(&mh).(*ListNode)
		
		// Push the next node of the popped node into the min heap.
		// TC: O((N-k) log k) as remaining element lists is being pushed in heap, and pushing take log K
		if node.Next != nil {
			heap.Push(&mh, node.Next)
		}
		
		node.Next = nil
		
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	
	return head
}

func main() {
	// Test cases
	ll := mergeKLists([]*ListNode{
		nil,
		{1, &ListNode{2, nil}},
	})
	for ll != nil {
		fmt.Print(ll.Val)
		ll = ll.Next
	}
	fmt.Println("")
	
	ll1 := mergeKLists([]*ListNode{
		nil,
		{1, nil},
	})
	for ll1 != nil {
		fmt.Print(ll1.Val)
		ll1 = ll1.Next
	}
	fmt.Println("")
	
	ll2 := mergeKLists([]*ListNode{
		{1, nil},
	})
	for ll2 != nil {
		fmt.Print(ll2.Val)
		ll2 = ll2.Next
	}
	fmt.Println("")
}