package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	hp, lp, c := 0, 0, 1
	for hp < len(nums) {
		if nums[hp] != nums[lp] {
			nums[lp+1] = nums[hp]
			lp++
			c++
		}
		hp++
	}
	fmt.Println(nums)
	return c
}

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,4,4}))
}
