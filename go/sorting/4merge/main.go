package main

import "fmt"

// Draw Choice diagram

/*
Intuition Behind Merge Sort:

Merge sort uses a "divide and conquer" strategy, which is often used in computer science to solve complex problems. Here's the breakdown:

Divide:

Imagine you have a pile of unsorted cards.
You repeatedly divide the pile into two roughly equal sub-piles until you have individual cards or very small sub-piles (which are inherently sorted).
Conquer (Sort):

At the smallest level, a single card or a pile of two cards is trivially sorted.
Merge:

Now, you start merging the sorted sub-piles back together.
You take two sorted sub-piles and create a new sorted pile by comparing the top cards of each sub-pile and repeatedly taking the smaller one.
You continue merging the sorted sub-piles until you have one final sorted pile.
In simpler terms:

You break down the problem into smaller, easier-to-solve subproblems (sorting small sub-piles).
Then, you combine the solutions to the subproblems to solve the original problem (merging the sorted sub-piles).
Why it's called "Merge Sort":

Because the core operation is "merging" sorted sub-arrays.
Key Idea:

It's a recursive algorithm.
It's efficient because merging two sorted arrays is a relatively fast operation.
It has a guaranteed time complexity of O(n log n) in all cases (best, average, and worst).
It's a stable sorting algorithm.
It is not an in-place sorting algorithm.
*/
func merge(arr []int) []int {
	// fmt.Println(arr)
	l := len(arr)
	if l <= 1 {
		return arr
	}
	mid := l / 2
	left := merge(arr[:mid])
	right := merge(arr[mid:])
	t:=mergeSort(left, right)
	// fmt.Println(t)
	return t
}

func mergeSort(left, right []int) []int {
	ans := []int{}
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			ans = append(ans, left[l])
			l++
		} else {
			ans = append(ans, right[r])
			r++
		}
	}
	for ; l < len(left); l++ {
		ans = append(ans, left[l])
	}
	for ; r < len(right); r++ {
		ans = append(ans, right[r])
	}
	return ans
}

func main() {
	fmt.Println(merge([]int{3, 1, 7, 2}))
	fmt.Println(merge([]int{3, 1, 7, 2, 9, 8}))
}
