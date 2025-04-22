package main

import "container/heap"

type Elements struct {
	val  int
	freq int
}

type MaxHeap []Elements

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Elements))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	numDict := make(map[int]int)
	for _, v := range nums {
		if _, ok := numDict[v]; ok {
			numDict[v]++
		} else {
			numDict[v] = 1
		}
	}
	h := MaxHeap{}
	heap.Init(&h)
	for k, f := range numDict {
		heap.Push(&h, Elements{
			val:  k,
			freq: f,
		})
	}
	res := []int{}
	if len(h) >= k {
		for range k {
			e := heap.Pop(&h).(Elements)
			res = append(res, e.val)
		}
	}
	return res
}

func topKFrequent1(nums []int, k int) []int {
	l := len(nums)
	numDict := make(map[int]int)
	for _, v := range nums {
		if _, ok := numDict[v]; ok {
			numDict[v]++
		} else {
			numDict[v] = 1
		}
	}
	// Length of t is equal to l because sum of occurrence of each element can't be more than total number of elements
	t := make([][]int, l)
	for val, f := range numDict {
		t[f-1] = append(t[f-1], val)
	}
	res := []int{}
	for i := l - 1; i >= 0; i-- {
		if t[i] != nil {
			res = append(res, t[i]...)
			if len(res) == k {
				break
			}
		}
	}
	return res
}
