package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numbers := make(map[int]any, len(nums))
	for _, num := range nums {
		if _, ok := numbers[num]; ok {
			return true
		}
		numbers[num] = num
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2}))
	fmt.Println(containsDuplicate([]int{1, 2, 1}))

	// Edge Case
	fmt.Println(containsDuplicate([]int{}))
}
