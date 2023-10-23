

import (
	"fmt"
	"math"
)

type MinHeap struct {
	items []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Insert(item int) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

func (h *MinHeap) ExtractMin() (int, bool) {
	if len(h.items) == 0 {
		return 0, false
	}

	min := h.items[0]
	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]
	h.heapifyDown(0)

	return min, true
}

func (h *MinHeap) heapifyUp(index int) {
	parentIndex := (index - 1) / 2

	for index > 0 && h.items[index] < h.items[parentIndex] {
		h.swap(index, parentIndex)
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (h *MinHeap) heapifyDown(index int) {
	minChildIndex := h.getMinChildIndex(index)

	for minChildIndex != -1 && h.items[index] > h.items[minChildIndex] {
		h.swap(index, minChildIndex)
		index = minChildIndex
		minChildIndex = h.getMinChildIndex(index)
	}
}

func (h *MinHeap) getMinChildIndex(index int) int {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	if leftChildIndex >= len(h.items) {
		return -1
	}

	if rightChildIndex >= len(h.items) || h.items[leftChildIndex] < h.items[rightChildIndex] {
		return leftChildIndex
	}

	return rightChildIndex
}

func (h *MinHeap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

func main() {
	heap := NewMinHeap()

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(15)
	heap.Insert(20)
	heap.Insert(7)

	min, ok := heap.ExtractMin()
	if ok {
		fmt.Println("Minimum element extracted:", min)
	}

	fmt.Println("Heap after extraction:", heap.items)
}