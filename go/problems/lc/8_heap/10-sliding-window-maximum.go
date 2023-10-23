// https://leetcode.com/problems/sliding-window-maximum/
// https://www.youtube.com/watch?v=LiSdD3ljCIE&list=PLEJXowNB4kPyP2PdMhOUlTY6GrRIITx28&index=16

package main

import "container/heap"

type element struct {
	value int
	index int
}

type MaxHeap []element

// Len returns the length of the MaxHeap.
func (h MaxHeap) Len() int {
	return len(h)
}

// Swap swaps the elements at indices i and j.
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less returns true if the element at index i is less than the element at index j.
func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

// Pop removes and returns the element at the top of the MaxHeap.
func (h *MaxHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

// Push adds an element x to the MaxHeap.
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

// ReadTop returns the element at the top of the MaxHeap.
func (h MaxHeap) ReadTop() element {
	return h[0]
}

// maxSlidingWindow returns the maximum elements in a sliding window of size k.
/*
The given function  `maxSlidingWindow`  takes in an array of integers  `nums`  and an integer  `k` . 
It returns a new array  `ans`  which contains the maximum element in each sliding window of size  `k`  as it moves through the input array.
The function first initializes an empty max heap  `mh`  and an empty array  `ans` . 
It then iterates through each element in the input array using a for loop.
Inside the loop, it pushes a new element into the max heap  `mh`  with the current element value and its index. 
This element represents the current window being processed.
After pushing the new element, the function checks if the length of the max heap  `mh`  is greater than  `k-1` . 
If so, it means that the window has moved beyond the size  `k`  and we need to pop elements from the max heap until the top element's index is within the current window.
If the top element's index is within the current window, it means that it is the maximum element for that window. 
The function appends the top element's value to the  `ans`  array and breaks out of the inner loop.
If the top element's index is not within the current window, it means that it is no longer part of the window and we can safely pop it from the max heap.
Finally, the function returns the  `ans`  array containing the maximum elements for each sliding window.

*/
func maxSlidingWindow(nums []int, k int) []int {
	mh := MaxHeap{}
	ans := []int{}
	for key, num := range nums {
		heap.Push(&mh, element{
			value: num,
			index: key,
		})
		for mh.Len() > k-1 {
			e := mh.ReadTop()
			if e.index <= key && e.index > key-k {
				ans = append(ans, e.value)
				break
			} else {
				heap.Pop(&mh)
			}
		}
	}
	return ans
}