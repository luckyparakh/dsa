package main

import "fmt"

func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if _, found := index[num]; found {
			return []int{i, index[num]}
		}
		index[diff] = i
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3,3}, 6))
	fmt.Println(twoSum([]int{1,3,3}, 6))
}
