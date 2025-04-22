// https://leetcode.com/problems/sort-colors/
// https://g.co/gemini/share/51fd29a097ba

package main

import "fmt"

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
	sortColorsBinary([]bool{false, true, true, false, false})
	sortColorsMultiple([]int{2, 0, 2, 1, 1, 0, 3, 0, 3})
}

func sortColors(nums []int) {
	driver(nums, 1)
}

func driver(nums []int, pivot int) {
	smaller, equal, larger := 0, 0, len(nums)-1

	for equal <= larger {
		if nums[equal] < pivot {
			nums[smaller], nums[equal] = nums[equal], nums[smaller]
			equal += 1
			smaller += 1
		} else if nums[equal] == pivot {
			equal += 1
		} else {
			nums[larger], nums[equal] = nums[equal], nums[larger]
			larger -= 1 // Think, Why not equal++, check LLM link for info
		}
	}
	fmt.Println((nums))
}

func sortColorsBinary(b []bool) {
	d := func(nums []bool, pivot bool) {
		equal, larger := 0, len(nums)-1

		for equal <= larger {
			if nums[equal] == pivot {
				equal += 1
			} else {
				nums[larger], nums[equal] = nums[equal], nums[larger]
				larger -= 1 // Think, Why not equal++, check LLM link for info
			}
		}
		fmt.Println((nums))
	}
	d(b, false)
}

// sortColorsMultiple sorts an array of integers containing multiple colors represented by integers.
// It uses a helper function driverM to partition the array based on the given color values.
//
// Parameters:
// - i: A slice of integers representing colors.
//
// The function initializes an index to 0 and defines a slice nums containing the color values to sort by.
// It then iterates over the nums slice, calling the driverM function to partition the array based on each color value.
// The index is updated to reflect the new partition point after each call to driverM.
// Finally, the function prints the sorted array.

func sortColorsMultiple(i []int) {
	// TC: n+n/2+n/4 --- =2n-1 = o(n)
	index := 0
	nums := []int{1, 2}
	for _,k := range nums {
		index = driverM(i[index:], k)
	}
	fmt.Println(i)
}

func driverM(nums []int, pivot int) int {
	fmt.Println(nums,pivot)
	smaller, equal, larger := 0, 0, len(nums)-1
	nextNumIndex := 0
	for equal <= larger {
		if nums[equal] < pivot {
			nums[smaller], nums[equal] = nums[equal], nums[smaller]
			equal += 1
			smaller += 1
		} else if nums[equal] == pivot {
			equal += 1
		} else {
			nums[larger], nums[equal] = nums[equal], nums[larger]
			larger -= 1 // Think, Why not equal++, check LLM link for info
		}
	}
	// Find next partition start index
	for k, v := range nums {
		if v > pivot+1 {
			nextNumIndex = k
			break
		}
	}
	// fmt.Println(nums, nextNumIndex)
	return nextNumIndex
}
