// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/triplet-sum-close-to-target-medium

package main

import (
	"math"
	"sort"
)

// Solution struct to encapsulate the functionality
type Solution struct{}

func (s *Solution) searchTriplet(arr []int, targetSum int) int {
	minDiff :=100_00_00
	minSum := 100_00_00

	// sort the array to use two pointer approach
	sort.Ints(arr)

	for i := 0; i < len(arr)-1; i++ {
		left := i + 1
		right := len(arr) - 1

		// two pointer approach
		for left < right {
			// calculate the sum of the triplet
			sum := arr[i] + arr[left] + arr[right]

			// check if we have found the exact target sum
			if sum == targetSum {
				return sum
			}

			// calculate the difference between the current sum and target sum
			diff := int(math.Abs(float64(targetSum - sum)))

			// update the minimum difference and corresponding sum if needed
			if diff < minDiff {
				minDiff = diff
				minSum = sum
			}

			// if the current difference is equal to the minimum difference found so far,
			if diff == minDiff {
				if sum < minSum {
					minSum = sum
				}
			}

			// move the pointers based on the comparison of sum and targetSum
			if sum < targetSum {
				left++
			}
			if sum > targetSum {
				right--
			}
		}
	}
	return minSum
}
