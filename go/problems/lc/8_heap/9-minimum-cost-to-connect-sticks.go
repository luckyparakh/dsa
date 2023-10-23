// https://github.com/doocs/leetcode/blob/main/solution/1100-1199/1167.Minimum%20Cost%20to%20Connect%20Sticks/README_EN.md

package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MinHeap) Pop() any {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func connectSticks(sticks []int) int {
	ans := 0
	if len(sticks) <= 1 {
		return 0
	}
	mh := MinHeap(sticks)
	heap.Init(&mh)
	for mh.Len() > 1 {
		sum := heap.Pop(&mh).(int) + heap.Pop(&mh).(int)
		ans += sum
		heap.Push(&mh, sum)
	}
	return ans
}

func main() {
	fmt.Println(connectSticks([]int{1, 8, 3, 5}))          //30
	fmt.Println(connectSticks([]int{2, 4, 3}))             //14
	fmt.Println(connectSticks([]int{8, 4, 6, 12}))         //58
	fmt.Println(connectSticks([]int{20, 4, 8, 2}))         //54
	fmt.Println(connectSticks([]int{1, 2, 5, 10, 35, 89})) //224
	fmt.Println(connectSticks([]int{2, 2, 3, 3}))          //20
	fmt.Println(connectSticks([]int{3}))                   //0
}
