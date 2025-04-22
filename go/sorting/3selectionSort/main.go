package main

import "fmt"

/*
Intuition Behind Selection Sort:

Imagine you have a pile of unsorted cards on a table, and you want to arrange them in ascending order.

Find the Smallest: You scan the entire pile to find the smallest card.
Place it First: You take that smallest card and place it at the beginning of a new, sorted pile.
Repeat: You repeat this process for the remaining unsorted cards:
Scan the remaining unsorted cards to find the next smallest.
Place it at the end of the sorted pile.
In simpler terms:

You're essentially "selecting" the smallest element from the unsorted part and moving it to its correct sorted position.
You build the sorted portion of the array one element at a time, always adding the smallest remaining element to the end.
Why it's called "Selection Sort":

Because you're repeatedly "selecting" the smallest element from the unsorted portion.

Key Idea:

The algorithm divides the array into two parts: a sorted portion (initially empty) and an unsorted portion (initially the entire array).
In each iteration, it finds the smallest element in the unsorted portion and moves it to the end of the sorted portion.

*/

func selectionSort(arr []int) {
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}
func main() {
	selectionSort([]int{3, 2, 4, 1})
}
