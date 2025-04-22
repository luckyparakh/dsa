package main

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	items []int
}

func (h *Heap) getParentIndex(i int) int {
	return (i - 1) / 2
}
func (h *Heap) getLeftChild(i int) int {
	return 2*i + 1
}
func (h *Heap) getRightChild(i int) int {
	return 2*i + 1
}
func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h *Heap) bottomToTop(index int) {
	pi := h.getParentIndex(index)
	for h.items[index] > h.items[pi] {
		h.swap(index, pi)
		index = pi
		pi = h.getParentIndex(index)
	}
}
func (h *Heap) topToBottom(index int) {
	lc := h.getLeftChild(index)
	rc := h.getRightChild(index)
	childCompare := -1
	l := len(h.items)
	for lc < l {
		if lc == l-1 {
			childCompare = lc
		}
		if h.items[lc] > h.items[rc] {
			childCompare = lc
		} else {
			childCompare = rc
		}
		if h.items[childCompare] > h.items[index] {
			h.swap(childCompare, index)
			index = childCompare
			lc = h.getLeftChild(index)
			rc = h.getRightChild(index)
		} else {
			return
		}
	}
}
func (h *Heap) push(val int) {
	h.items = append(h.items, val)
	h.bottomToTop(len(h.items) - 1)
}
func (h *Heap) pop() (int, bool) {
	if len(h.items) == 0 {
		return 0, false
	}
	l := len(h.items) - 1
	ans := h.items[0]
	h.items[0] = h.items[l]
	h.items = h.items[:l]
	h.topToBottom(0)
	return ans, true
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() any {
	old := (*h)
	l := len(old) - 1
	*h = old[:l]
	return old[l]
}
func main() {
	hp := Heap{}
	hp.push(10)
	hp.push(40)
	hp.push(5)
	if min, ok := hp.pop(); ok {
		fmt.Println(min)
	}

}
func lastStoneWeight(stones []int) int {
	h := maxHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		y, x := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if y > x {
			heap.Push(&h, y-x)
		}
	}
	return h[0]
}
