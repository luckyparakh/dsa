// https://leetcode.com/problems/find-k-closest-elements/

package main

import (
	"container/heap"
	"sort"
)

type element struct {
	key  int
	diff int
}

// MaxHeap is a custom type representing a heap of integers
type MaxHeap []element

// Len returns the length of the heap
func (h MaxHeap) Len() int {
	return len(h)
}

// Less returns whether the element at index i is less than the element at index j
func (h MaxHeap) Less(i, j int) bool {
	if h[i].diff == h[j].diff {
		return h[i].key > h[j].key
	}
	return h[i].diff > h[j].diff
}

// Swap swaps the elements at indexes i and j
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

// Pop removes and returns the element at the top of the heap
func (h *MaxHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}
func findClosestElements(arr []int, k int, x int) []int {
	mh := MaxHeap{}
	for _, a := range arr {
		heap.Push(&mh, element{
			diff: mod(a, x),
			key:  a,
		})
		if mh.Len() > k {
			heap.Pop(&mh)
		}
	}
	ans := []int{}
	for mh.Len() > 0 {
		ans = append(ans, heap.Pop(&mh).(element).key)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func mod(a, x int) int {
	diff := a - x
	if diff < 0 {
		return -diff
	}
	return diff
}
