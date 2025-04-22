// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

import "fmt"

// removeDuplicates removes duplicates from a sorted slice of integers in-place
// and returns the length of the slice after duplicates have been removed.
// It also prints the maximum length of any repeating number sequence found in the input slice.
//
// Parameters:
// - nums: A slice of integers sorted in non-decreasing order.
//
// Returns:
// - An integer representing the length of the slice after removing duplicates.
//
// Example:
// nums := []int{1, 1, 2, 2, 2, 3, 4, 4}
// length := removeDuplicates(nums)
// fmt.Println(length) // Output: 4
// fmt.Println(nums[:length]) // Output: [1, 2, 3, 4]
func removeDuplicates(nums []int) int {
	lp := 0
	s := 0
	maxLenOfRepeatingNumber := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			if i-s > maxLenOfRepeatingNumber {
				maxLenOfRepeatingNumber = i-s
			}
			s=i
			
			lp++
			nums[lp] = nums[i]
		}
	}
	fmt.Print(maxLenOfRepeatingNumber)
	return lp + 1
}
