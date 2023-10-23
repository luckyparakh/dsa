// https://leetcode.com/problems/sort-array-by-increasing-frequency/
package main

import "container/heap"

type element struct {
	value int
	count int
}
type MinHeap []element

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].value > h[j].value
	}
	return h[i].count < h[j].count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() any {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}
func (h *MinHeap) Push(e any) {
	*h = append(*h, e.(element))
}
func frequencySort(nums []int) []int {
	mh := MinHeap{}
	numCount := count(nums)
	for key, numCount := range numCount {
		heap.Push(&mh, element{
			value: key,
			count: numCount,
		})
	}
	ans := []int{}
	for len(mh) > 0 {
		v := heap.Pop(&mh).(element)
		for i := 0; i < v.count; i++ {
			ans = append(ans, v.value)
		}
	}
	return ans
}

// count returns a map of element counts in nums
func count(nums []int) map[int]int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	return mp
}
