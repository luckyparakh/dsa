// Problem: https://leetcode.com/problems/find-median-from-data-stream/
// Video Explanation: https://www.youtube.com/watch?v=EcNbRjEcb14

package main

import "fmt"

type MinHeap []int

// Pop removes and returns the minimum element from the heap
func (minH *MinHeap) Pop() any {
	if len(*minH) == 0 {
		return nil
	}
	v := (*minH)[0]
	l := len(*minH) - 1
	(*minH)[0] = (*minH)[l]
	*minH = (*minH)[:l]
	minH.heapifyBottom(0)
	return v
}

// heapifyBottom performs heapify operation on the heap from bottom to top
func (minH *MinHeap) heapifyBottom(pi int) {
	li, ri := leftChildIndex(pi), rightChildIndex(pi)
	childToCompareIndex := -1
	for li <= len(*minH)-1 {
		if li == len(*minH)-1 {
			childToCompareIndex = li
		} else if (*minH)[li] < (*minH)[ri] {
			childToCompareIndex = li
		} else {
			childToCompareIndex = ri
		}
		if (*minH)[childToCompareIndex] < (*minH)[pi] {
			minH.swap(childToCompareIndex, pi)
			pi = childToCompareIndex
			li, ri = leftChildIndex(pi), rightChildIndex(pi)
		} else {
			return
		}
	}
}

// Push adds a new element to the heap
func (minH *MinHeap) Push(v any) {
	*minH = append(*minH, v.(int))
	ci := len(*minH) - 1
	pi := parentIndex(ci)
	for (*minH)[ci] < (*minH)[pi] {
		minH.swap(pi, ci)
		ci = pi
		pi = parentIndex(ci)
	}
}

// swap swaps the elements at the given indices in the heap
func (minH *MinHeap) swap(i, j int) {
	(*minH)[i], (*minH)[j] = (*minH)[j], (*minH)[i]
}

// top returns the minimum element from the heap
func (minH *MinHeap) top() int {
	return (*minH)[0]
}

type MaxHeap []int

// Pop removes and returns the maximum element from the heap
func (maxH *MaxHeap) Pop() any {
	if len(*maxH) == 0 {
		return nil
	}
	v := (*maxH)[0]
	l := len(*maxH) - 1
	(*maxH)[0] = (*maxH)[l]
	*maxH = (*maxH)[:l]
	maxH.heapifyBottom(0)
	return v
}

// heapifyBottom performs heapify operation on the heap from bottom to top
func (maxH *MaxHeap) heapifyBottom(pi int) {
	li, ri := leftChildIndex(pi), rightChildIndex(pi)
	childToCompareIndex := -1
	for li <= len(*maxH)-1 {
		if li == len(*maxH)-1 {
			childToCompareIndex = li
		} else if (*maxH)[li] > (*maxH)[ri] {
			childToCompareIndex = li
		} else {
			childToCompareIndex = ri
		}
		if (*maxH)[childToCompareIndex] > (*maxH)[pi] {
			maxH.swap(childToCompareIndex, pi)
			pi = childToCompareIndex
			li, ri = leftChildIndex(pi), rightChildIndex(pi)
		} else {
			return
		}
	}
}

// Push adds a new element to the heap
func (maxH *MaxHeap) Push(v any) {
	*maxH = append(*maxH, v.(int))
	ci := len(*maxH) - 1
	pi := parentIndex(ci)
	for (*maxH)[ci] > (*maxH)[pi] {
		maxH.swap(pi, ci)
		ci = pi
		pi = parentIndex(ci)
	}
}

// swap swaps the elements at the given indices in the heap
func (maxH *MaxHeap) swap(i, j int) {
	(*maxH)[i], (*maxH)[j] = (*maxH)[j], (*maxH)[i]
}

// top returns the maximum element from the heap
func (maxH *MaxHeap) top() int {
	return (*maxH)[0]
}

// parentIndex returns the index of the parent node given the child index
func parentIndex(ci int) int {
	return (ci - 1) / 2
}

// leftChildIndex returns the index of the left child node given the parent index
func leftChildIndex(pi int) int {
	return pi*2 + 1
}

// rightChildIndex returns the index of the right child node given the parent index
func rightChildIndex(pi int) int {
	return pi*2 + 2
}

type MedianFinder struct {
	minH MinHeap
	maxH MaxHeap
}

// Constructor initializes a new MedianFinder
func Constructor() MedianFinder {
	return MedianFinder{
		minH: MinHeap{},
		maxH: MaxHeap{},
	}
}

// AddNum adds a new number to the MedianFinder
func (this *MedianFinder) AddNum(num int) {
	if len(this.maxH) == 0 || this.maxH.top() > num {
		this.maxH.Push(num)
	} else {
		this.minH.Push(num)
	}
	if len(this.maxH)-len(this.minH) > 1 {
		this.minH.Push(this.maxH.Pop())
	}
	if len(this.minH)-len(this.maxH) > 1 {
		this.maxH.Push(this.minH.Pop())
	}
}

// FindMedian returns the median of all the numbers added so far
func (this *MedianFinder) FindMedian() float64 {
	maxL, minL := len(this.maxH), len(this.minH)
	maxT := float64(this.maxH.top())
	if minL == 0 {
		return maxT
	}
	minT := float64(this.minH.top())
	if maxL > minL {
		return maxT
	} else if minL > maxL {
		return minT
	}
	return (maxT + minT) / 2
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}