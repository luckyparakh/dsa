// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
// https://www.youtube.com/watch?v=4WmTRFZilj8&list=PL_z_8CaSLPWeYfhtuKHj-9MpYb6XQJ_f2&index=7

package main

import "fmt"

// Minimum number will have feature that its both neighbor will be higher in value
// Mid will always divide array in two part out of one will always be unsorted.
// Unsorted array will have minimum number.

func findMin(nums []int) int {
	l := len(nums)
	s, e := 0, l-1
	for s <= e {
		mid := s + (e-s)/2
		next := (mid + 1) % l
		prev := (mid + l - 1) % l
		if nums[mid] < nums[next] && nums[mid] < nums[prev] {
			return nums[mid]
		}
		if nums[mid] > nums[l-1] {
			s = mid + 1
		} else if nums[mid] < nums[0] {
			e = mid - 1
		} else {
			// Array is not rotated
			return nums[0]
		}
	}
	// Will never be called but func need return
	return -1
}
func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{1}))
	fmt.Println(findMin([]int{1, 2, 3, 4}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
