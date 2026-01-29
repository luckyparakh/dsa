// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/triplets-with-smaller-sum-medium

package main

import "sort"

// Solution struct to simulate a class
type Solution struct{}

func (s *Solution) searchTriplets(arr []int, target int) int {
	count := 0
	// sort the array to use two pointer approach
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			sum := arr[left] + arr[right] + arr[i]
			if sum < target {
				// KEY INSIGHT: All pairs between left and right work!
				// Why ? Because the array is sorted.
				// So, if we have found arr[left] + arr[right] + arr[i] < target
				// then, all elements between left and right will also work with arr[i] 
				// as now arr[left] and arr[i] are fixed, and if we decrease arr[right],
				// the sum will only decrease further, so all rigths with left and i are valid pairs.
				// For example, if we have arr = [1, 2, 3, 4, 5], target = 8
				// and we are at i=0 (arr[i]=1), left=1 (arr[left]=2), right=4 (arr[right]=5)
				// then arr[i] + arr[left] + arr[right] = 1 + 2 + 5 = 8 (not valid)
				// but if we move right to 3 (arr[right]=4), then arr[i] + arr[left] + arr[right] = 1 + 2 + 4 = 7 (valid)
				// and also arr[i] + arr[left] + arr[2] = 1 + 2 + 3 = 6 (valid)
				// Hence, we can count all those pairs in one go.
				count += right - left
				left++
			} else {
				// sum>=target, need smaller sum
				right--
			}
		}
	}
	return count
}
