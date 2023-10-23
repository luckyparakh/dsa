// https://leetcode.com/problems/k-closest-points-to-origin/
package main

import "container/heap"

// MaxHeap is a custom type representing a heap of distances
type MaxHeap [][]int

// Len returns the length of the heap
func (h MaxHeap) Len() int {
	return len(h)
}

// Less returns whether the distance at index i is less than the distance at index j
func (h MaxHeap) Less(i, j int) bool {
	return distance(h[i]) > distance(h[j])
}

// Swap swaps the distances at indexes i and j
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a distance to the heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// Pop removes and returns the distance at the top of the heap
func (h *MaxHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

// kClosest returns the k closest points to the origin
func kClosest(points [][]int, k int) [][]int {
	mh := MaxHeap{}
	for i := 0; i < len(points); i++ {
		heap.Push(&mh, points[i])
		if mh.Len() > k {
			heap.Pop(&mh)
		}
	}
	return mh
}

// distance calculates the distances from the origin to each point
func distance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}