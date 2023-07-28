// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

package main

import "fmt"

// func findLast(n []int, num int) int {
// 	s := 0
// 	e := len(n) - 1
// 	result := -1
// 	for s <= e {
// 		mid := s + (e-s)/2
// 		if n[mid] == num {
// 			result = mid
// 			s = mid + 1
// 		} else if n[mid] > num {
// 			e = mid - 1
// 		} else {
// 			s = mid + 1
// 		}
// 	}
// 	return result
// }

func findFirst(n []int, num int, first bool) int {
	s := 0
	e := len(n) - 1
	result := -1
	for s <= e {
		mid := s + (e-s)/2
		if n[mid] == num {
			result = mid
			if first {
				e = mid - 1
			} else {
				s = mid + 1
			}

		} else if n[mid] > num {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return result
}

func searchRange(nums []int, target int) []int {
	op := []int{}
	op = append(op, findFirst(nums, target, true))
	op = append(op, findFirst(nums, target, false))
	// To count element occurrence calculate op[1]-op[0]+1
	return op
}
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 6))
}
