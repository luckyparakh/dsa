package main

import "fmt"

// https://www.youtube.com/watch?v=3DYIgTC4T1o&t=917s

type MaxHeap struct {
	items []int
}

func NewMaxHeap() *MaxHeap { return &MaxHeap{} }

func (this *MaxHeap) Insert(item int) {
	// insert the item at end of heap and heapify the heap
	this.items = append(this.items, item)
	this.heapifyBottomToTop(len(this.items) - 1)
}

// Extracts largest element from the heap and remove it from heap
func (this *MaxHeap) Extract() int {
	if len(this.items) == 0 {
		return -1
	}
	firstItem := this.items[0]
	lastIndex := len(this.items) - 1
	this.items[0] = this.items[lastIndex]
	this.items = this.items[:lastIndex]
	this.heapifyTopToBottom(0)
	return firstItem
}

func (this *MaxHeap) heapifyBottomToTop(index int) {
	pIndex := this.getParentIndex(index)
	for this.items[index] > this.items[pIndex] {
		this.swap(index, pIndex)
		index = pIndex
		pIndex = this.getParentIndex(index)
	}
}

func (this *MaxHeap) heapifyTopToBottom(index int) {
	rightChildIndex := this.getRightChildIndex(index)
	leftChildIndex := this.getLeftChildIndex(index)
	childToCompare := -1
	l := len(this.items) - 1
	for leftChildIndex <= l {
		if leftChildIndex == l {
			childToCompare = leftChildIndex
		} else if this.items[leftChildIndex] > this.items[rightChildIndex] {
			childToCompare = leftChildIndex
		} else {
			childToCompare = rightChildIndex
		}
		if this.items[childToCompare] > this.items[index] {
			this.swap(childToCompare, index)
			index = childToCompare
			rightChildIndex = this.getRightChildIndex(index)
			leftChildIndex = this.getLeftChildIndex(index)
		} else {
			return
		}
	}
}

func (this *MaxHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}
func (this *MaxHeap) getRightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}
func (this *MaxHeap) getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}
func (this *MaxHeap) swap(index1, index2 int) {
	this.items[index1], this.items[index2] = this.items[index2], this.items[index1]
}
func main() {
	m := NewMaxHeap()
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	// buildHeap := []int{10, 20, 30, 5, 15, 17,19}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	fmt.Println("Extract")
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
