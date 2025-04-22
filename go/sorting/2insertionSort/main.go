package main

import "fmt"
/*
Intuition Behind Insertion Sort:

Imagine you're organizing a hand of playing cards.

Start with the Second Card: You pick up the second card in your hand.

Compare and Shift: You compare it with the card(s) already in your hand (which are considered sorted). If the new card is smaller, you shift the larger cards to the right to make space for the new card.

Insert: You then insert the new card into its correct position.

Repeat: You repeat this process for each remaining card, always comparing and shifting within the already sorted portion of your hand.

In simpler terms:

You're building a sorted portion of your hand one card at a time.
Each new card is "inserted" into its proper place within the already sorted cards.
Why it's called "Insertion Sort":

Because you're repeatedly "inserting" a new element into its correct position within the sorted portion of the array.
Key Idea:

The algorithm divides the array into two parts: a sorted portion (initially containing the first element) and an unsorted portion (the rest of the array).
In each iteration, it takes the next element from the unsorted portion and inserts it into its correct position within the sorted portion.

*/
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
	fmt.Println(arr)
}
func main() {
	insertSort([]int{3, 2, 5, 1})
}
