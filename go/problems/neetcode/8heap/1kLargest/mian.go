package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	h minHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	// tmp := minHeap(nums)
	kl := KthLargest{
		k: k,
		h: minHeap(nums),
	}
	heap.Init(&kl.h)
	for kl.h.Len() > kl.k {
		heap.Pop(&kl.h)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h,val)
	fmt.Println(this.h)
	for this.h.Len() > this.k {
		heap.Pop(&this.h)
	}
	fmt.Println(this.h)
	return this.h[0]
}

type minHeap []int

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Len() int           { return len(h) }
func (h *minHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *minHeap) Pop() any {
	old := *h
	l := len(old) - 1
	*h = old[:l]
	return old[l]
}

func main() {
	kth := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kth.Add(3))
}
