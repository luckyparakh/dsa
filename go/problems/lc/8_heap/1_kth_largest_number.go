// Package main implements the solution to the problem "Kth Largest Element in a Stream" on LeetCode.
// The problem can be found at https://leetcode.com/problems/kth-largest-element-in-a-stream/.
// The solution uses a min heap to maintain the k largest elements in the stream.

package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is a type alias for a slice of integers.
type IntHeap []int

// Len returns the length of the heap.
func (h IntHeap) Len() int {
	return len(h)
}

// Less returns true if the element at index i is less than the element at index j.
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements at indexes i and j.
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the smallest element from the heap.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// KthLargest is a struct that represents the kth largest element in a stream.
type KthLargest struct {
	minHeap *IntHeap
	k       int
}

// Constructor creates a new instance of KthLargest.
func Constructor(k int, nums []int) KthLargest {
	// Type cast nums in to IntHeap which is a slice of int type
	tmp := IntHeap(nums)
	this := KthLargest{
		k:       k,
		minHeap: &tmp,
	}
	heap.Init(this.minHeap)
	for len(*this.minHeap) > this.k {
		heap.Pop(this.minHeap)
	}
	return this
}

// Add adds a value to the stream and returns the kth largest element.
func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	for len(*this.minHeap) > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

func main() {
	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	fmt.Println(k.Add(9))
	fmt.Println(k.Add(4))
}