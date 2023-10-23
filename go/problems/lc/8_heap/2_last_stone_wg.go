// https://leetcode.com/problems/last-stone-weight/

package main

import "container/heap"
// IntHeap is a custom type representing a heap of integers
type IntHeap []int

// Len returns the length of the heap
func (h IntHeap) Len() int { 
	return len(h) 
}

// Max Heap
// Less returns whether the element at index i is less than the element at index j
func (h IntHeap) Less(i, j int) bool { 
	return h[i] > h[j] 
}

// Swap swaps the elements at indexes i and j
func (h IntHeap) Swap(i, j int) { 
	h[i], h[j] = h[j], h[i] 
}

// Push adds an element to the heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the element at the top of the heap
func (h *IntHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

// lastStoneWeight calculates the weight of the last stone remaining after smashing the stones together
func lastStoneWeight(stones []int) int {
	// Create a new IntHeap using the input stones
	h := IntHeap(stones)
	// Initialize the heap
	heap.Init(&h)
	// Continue smashing stones until there is only one stone left
	for h.Len() > 1 {
		// Remove the two largest stones from the heap
		large, small := heap.Pop(&h), heap.Pop(&h)
		// If the two stones are not equal, calculate the difference and add it back to the heap
		if large != small {
			heap.Push(&h, (large.(int) - small.(int)))
		}
	}
	// If there is one stone left, return its weight
	if h.Len() == 1 {
		return h[0]
	}
	// If there are no stones left, return 0
	return 0
}