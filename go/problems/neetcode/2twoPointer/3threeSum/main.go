package main

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	// tc: o(n2),sc:o(1)
	slices.Sort(nums)
	ans := [][]int{}
	for i := range nums {
		// If the current number nums[i] is greater than zero, because the array is sorted, any combination involving this and subsequent numbers will be greater than zero.
		// Since no sum with a positive total can equal zero, we can break out of the loop early.
		if nums[i] > 0 {
			break
		}
		// If the current number is the same as the previous number (and i is not 0),
		// we use continue to skip to the next iteration to avoid checking for duplicates.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		s, e := i+1, len(nums)-1
		for s < e {
			currSum := nums[s] + nums[e] + nums[i]
			if currSum == 0 {
				ans = append(ans, []int{nums[i], nums[s], nums[e]})
				s++
				e--
				// Consider this sorted arr:=[-10 -5 -5  0 0 5 5]
				// Let's say i is at index 1 (-5) so one of the combination is [-5,0,5] index (1,3,6) another 
				// can be elements at index (1,4,5) but this is duplicate, hence to remove this below code is needed.
				// Skip duplicate elements by moving the left pointer to the right.
				for s < e && nums[s] == nums[s-1] {
					s++
				}
				// Skip duplicate elements by moving the right pointer to the left.
				for s < e && nums[e] == nums[e+1] {
					e--
				}
			}
			if currSum > 0 {
				e--
			}
			if currSum < 0 {
				s++
			}
		}
	}
	return ans
}
