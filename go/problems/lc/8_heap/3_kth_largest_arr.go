// https://leetcode.com/problems/kth-largest-element-in-an-array/

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

func findKthLargest(nums []int, k int) int {
    h:=IntHeap(nums)
	heap.Init(&h)
	for i:=0;i<k-1;i++{
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

import (
    "container/heap"
    "testing"
)

func TestFindKthLargest(t *testing.T) {
    // Positive test case
    nums := []int{3, 6, 2, 1, 8, 4}
    k := 3
    expected := 4
    result := findKthLargest(nums, k)
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }

    // Negative test case
    nums = []int{}
    k = 1
    expected = 0
    result = findKthLargest(nums, k)
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }

    // Boundary test case
    nums = []int{5}
    k = 1
    expected = 5
    result = findKthLargest(nums, k)
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }

    // Edge test case
    nums = []int{2, 2, 2, 2, 2}
    k = 1
    expected = 2
    result = findKthLargest(nums, k)
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }
}