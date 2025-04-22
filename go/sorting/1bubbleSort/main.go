package main

import "fmt"
/*
**Intuition Behind Bubble Sort:**

Imagine you have a vertical tube filled with bubbles of different sizes, where the "size" represents the value of an element in an array.

1.  **Compare and Swap:** You start from the bottom of the tube and compare adjacent bubbles.
2.  **"Bubble Up" the Larger:** If a bubble is larger than the one above it, you swap their positions, effectively letting the larger bubble "bubble up" towards the top.
3.  **Repeat:** You repeat this process for each pair of adjacent bubbles, moving upwards through the tube.
4.  **Repeat Passes:** After one pass through the tube, the largest bubble will have reached the top. You repeat this process for the remaining bubbles, excluding the ones that have already "bubbled up" to their correct positions.

**In simpler terms:**

* You're repeatedly comparing adjacent elements and swapping them if they're in the wrong order.
* Larger elements "bubble up" towards their correct positions at the end of the array.

**Why it's called "Bubble Sort":**

* Because the larger elements "bubble up" to their correct positions like bubbles rising in water.

**Key Idea:**

* The algorithm repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
* The pass through the list is repeated until the list is sorted.
* Each pass moves the largest unsorted element to its correct position at the end of the unsorted part of the array.

*/
func bubbleSort(arr []int) {
	for i := range arr {
		swapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println(arr)
}

func main() {
	bubbleSort([]int{5, 3, 4, 1})
	bubbleSort([]int{1, 5, 4, 3})
}
