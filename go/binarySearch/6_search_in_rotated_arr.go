// https://leetcode.com/problems/search-in-rotated-sorted-array/
package main

import "fmt"

func findMin(nums []int) int {
	l := len(nums)
	s, e := 0, l-1
	for s <= e {
		m := s + (e-s)/2
		next := (m + 1) % l
		prev := (m + l - 1) % l

		if nums[m] < nums[next] && nums[m] < nums[prev] {
			return m
		}
		if nums[m] > nums[l-1] {
			s = m + 1
		} else if nums[m] < nums[0] {
			e = m - 1
		} else {
			return 0
		}

	}
	return -1
}
func BS(nums []int, target, s, e int) int {
	for s <= e {
		m := s + (e-s)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return -1
}
func search(nums []int, target int) int {
	minIndex := findMin(nums)
	index1 := BS(nums, target, 0, minIndex-1)
	if index1 == -1 {
		index2 := BS(nums, target, minIndex, len(nums)-1)
		if index2 == -1 {
			return -1
		}
		return index2
	}
	return index1
}
func main() {
	// fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	// fmt.Println(findMin([]int{1}))
	// fmt.Println(findMin([]int{1, 2, 3, 4}))
	// fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))

	fmt.Println(search([]int{3, 4, 5, 1, 2}, 5))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{1, 2, 3, 4}, 2))
}
