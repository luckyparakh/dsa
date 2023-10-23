// https://leetcode.com/problems/top-k-frequent-elements/
package main

import (
	"container/heap"
)

type element struct {
	value int
	count int
}

// MinHeap is a custom type representing a heap of integers
type MinHeap []element

// Len returns the length of the heap
func (h MinHeap) Len() int {
	return len(h)
}

// Less returns whether the element at index i is less than the element at index j
func (h MinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

// Swap swaps the elements at indexes i and j
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

// Pop removes and returns the element at the top of the heap
func (h *MinHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}
/*
This function takes an array of integers and an integer k as input. It returns the top k frequent elements from the input array.
First, an empty array ans is created to store the result. The function then calls the count function to count the frequency of each number in the input array. 
The count function returns a map where the keys are the numbers and the values are their frequencies.
A min heap called minHeap is created to store the elements with their frequencies. The heap is initially empty. 
The function then iterates through the map returned by the count function. 
For each element in the map, a new element is created with the number as the value and the frequency as the count. 
This element is then pushed into the min heap.
If the size of the min heap exceeds k, the element with the smallest frequency is removed from the heap using the heap.Pop function. 
This ensures that the heap only contains the top k frequent elements at any given time.
Finally, the function iterates through the elements in the min heap and appends their values to the ans array. 
This array is then returned as the result.
*/
func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	numCount := count(nums)
	minHeap := MinHeap{}

	// Build the min heap
	for num, count := range numCount {
		heap.Push(&minHeap, element{
			value: num,
			count: count,
		})
		if minHeap.Len() > k {
			heap.Pop(&minHeap)
		}
	}

	// Extract the top k frequent elements from the min heap
	for _, e := range minHeap {
		ans = append(ans, e.value)
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
