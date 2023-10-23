// https://leetcode.com/problems/sort-characters-by-frequency/
package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type element struct {
	value rune
	count int
}

type MaxHeap []element

// Len returns the length of the MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Less compares two elements in the MaxHeap and returns true if the element at index i has a higher count than the element at index j
func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

// Swap swaps two elements in the MaxHeap
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Pop removes and returns the element with the highest count from the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

// Push adds an element to the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

// frequencySort sorts the characters in the string s by their frequency and returns the sorted string
func frequencySort(s string) string {
	mh := MaxHeap{}
	sCount := count(s)
	for key, numCount := range sCount {
		heap.Push(&mh, element{
			value: key,
			count: numCount,
		})
	}

	sb := strings.Builder{}
	for len(mh) > 0 {
		v := heap.Pop(&mh).(element)
		for i := 0; i < v.count; i++ {
			sb.WriteRune(v.value)
		}
	}

	return sb.String()
}

// count returns a map of element counts in the string s
func count(s string) map[rune]int {
	mp := map[rune]int{}
	for _, char := range s {
		mp[char]++
	}
	return mp
}

func main() {
	fmt.Println(frequencySort("tree"))
}
