// https://leetcode.com/problems/k-closest-points-to-origin/
package main

import (
	"container/heap"
	"math"
)

// distance represents the distance from the origin to a point
type distance struct {
	d   float64 // distance
	key int     // index of the point
}

// DHeap is a custom type representing a heap of distances
type DHeap []distance

// Len returns the length of the heap
func (h DHeap) Len() int {
	return len(h)
}

// Less returns whether the distance at index i is less than the distance at index j
func (h DHeap) Less(i, j int) bool {
	return h[i].d < h[j].d
}

// Swap swaps the distances at indexes i and j
func (h DHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a distance to the heap
func (h *DHeap) Push(x interface{}) {
	*h = append(*h, x.(distance))
}

// Pop removes and returns the distance at the top of the heap
func (h *DHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

// kClosest returns the k closest points to the origin
func kClosest(points [][]int, k int) [][]int {
	distanceSlice := getDistance(points)
	dh := DHeap(distanceSlice)
	heap.Init(&dh)
	ans := [][]int{}
	for i := 1; i <= k; i++ {
		ans = append(ans, points[heap.Pop(&dh).(distance).key])
	}
	return ans
}

// getDistance calculates the distances from the origin to each point
func getDistance(points [][]int) []distance {
	ans := []distance{}
	for k, point := range points {
		ans = append(ans, distance{
			d:   math.Sqrt(float64(point[0]*point[0] + point[1]*point[1])),
			key: k,
		})
	}
	return ans
}
