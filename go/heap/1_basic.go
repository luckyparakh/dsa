package main

import "fmt"

// MaxHeap is a struct representing a max heap data structure.
type MaxHeap struct {
	items []int
}

// NewMaxHeap creates a new instance of MaxHeap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Insert inserts an item into the max heap.
func (this *MaxHeap) Insert(item int) {
	this.items = append(this.items, item)
	this.heapifyBottomToTop(len(this.items) - 1)
}

// Extract extracts the largest element from the heap and removes it from the heap.
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

// heapifyBottomToTop adjusts the heap from bottom to top to maintain the max heap property.
func (this *MaxHeap) heapifyBottomToTop(index int) {
	pIndex := this.getParentIndex(index)
	for this.items[index] > this.items[pIndex] {
		this.swap(index, pIndex)
		index = pIndex
		pIndex = this.getParentIndex(index)
	}
}

// heapifyTopToBottom adjusts the heap from top to bottom to maintain the max heap property.
/*
This function  `heapifyTopToBottom`  is a method of the  `MaxHeap`  struct. 
It is responsible for maintaining the max heap property by comparing the parent node with its children and swapping them if necessary.
The function takes an  `index`  parameter, which represents the index of the parent node. It begins by calculating the indices of the left and right child nodes using helper functions  `getLeftChildIndex`  and  `getRightChildIndex` .
The function then enters a loop that continues until the left child index is less than or equal to the length of the heap. 
This loop iterates through the heap from top to bottom, comparing the parent node with its children and swapping them if necessary.
Inside the loop, the function checks if the left child index is equal to the length of the heap. 
If it is, it means that the parent node only has a left child and no right child. In this case, the left child is assigned to the  `childToCompare`  variable.
If the parent node has both a left and right child, the function compares the values of the left and right child nodes. 
The index of the child with the larger value is assigned to the  `childToCompare`  variable.
Next, the function checks if the value of the child node to compare is greater than the value of the parent node. 
If it is, it means that the child node violates the max heap property, so a swap is performed between the parent and child nodes. The  `swap`  method of the  `MaxHeap`  struct is called to perform the swap. After the swap, the index is updated to the index of the swapped child node.
Finally, the function recalculates the indices of the left and right child nodes using the updated index. 
This process continues until the parent node is in its correct position, or until the loop terminates if the parent node is already greater than both of its children.
The purpose of this function is to maintain the max heap property by ensuring that the parent node is always greater than or equal to its children. This is necessary for efficient operations on a max heap, such as inserting and extracting the maximum element.
*/
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

// getParentIndex returns the index of the parent node for a given child index.
func (this *MaxHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// getRightChildIndex returns the index of the right child node for a given parent index.
func (this *MaxHeap) getRightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}

// getLeftChildIndex returns the index of the left child node for a given parent index.
func (this *MaxHeap) getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}

// swap swaps the positions of two items in the heap.
func (this *MaxHeap) swap(index1, index2 int) {
	this.items[index1], this.items[index2] = this.items[index2], this.items[index1]
}

func main() {
	m := NewMaxHeap()
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	// buildHeap := []int{10, 20, 30, 5, 15, 17, 19}
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