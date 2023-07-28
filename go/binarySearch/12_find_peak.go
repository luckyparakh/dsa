package main
// Find 1st peak in array
// Sub case of it is find peak in mountain
// https://leetcode.com/problems/peak-index-in-a-mountain-array/submissions/
import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// if len(nums) == 2 {
	// 	if nums[0] > nums[1] {
	// 		return 0
	// 	}
	// 	return 1
	// }
	s := 0
	e := len(nums) - 1

	for s <= e {
		m := s + (e-s)/2
		
		// Edge Case when m=0
		if m == 0 {
			if nums[m] > nums[m+1] {
				return 0
			} else {
				return 1
			}
		// Edge case when m=len-1
		}else if m == len(nums)-1 {
			if nums[m-1] < nums[m] {
				return len(nums) - 1
			} else {
				return len(nums) - 2
			}
		}else if nums[m-1] < nums[m] && nums[m] > nums[m+1] {
			return m
		}else if nums[m] < nums[m+1] {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return -1
}
func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1, 2, 3, 4}))
	fmt.Println(findPeakElement([]int{4, 3, 2, 1, 0}))
	fmt.Println(findPeakElement([]int{4, 3}))
	fmt.Println(findPeakElement([]int{4, 30}))
	fmt.Println(findPeakElement([]int{30}))
	fmt.Println(findPeakElement([]int{3, 0, 2}))
}
