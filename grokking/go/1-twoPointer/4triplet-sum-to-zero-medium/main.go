// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/triplet-sum-to-zero-medium
package main

import (
	"sort"
)

type Solution struct{}

func (s *Solution) searchTriplets(arr []int) [][]int {
	triplets := make([][]int, 0)
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		// We skip duplicates for the first element of the triplet
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		// Since the array is sorted, if the current number is greater than 0,
		// we cannot find any triplet that sums to zero
		// Why not >= 0? Because we can have triplets like [0, 0, 0]
		if arr[i] > 0 {
			break
		}

		// now we need to find a pair in arr[i+1..end] which sums to -arr[i]
		l, r := i+1, len(arr)-1
		ts := arr[i] * -1

		// two pointer approach to find the pair
		for l < r {
			s := arr[l] + arr[r]
			if s == ts {
				// current i, l, r form a triplet
				triplets = append(triplets, []int{arr[i], arr[l], arr[r]})
				r--
				l++

				// skip duplicates for the second and third elements of the triplet
				// Why l< r? Because we need to ensure that the two pointers do not cross each other while skipping duplicates.
				for l < r && arr[r] == arr[r+1] {
					// check if r is a duplicate of the previous element is r+1.
					// Why previous element is r+1? Because we have already moved r one step to the left.
					r--
				}
				for l < r {
					// check if l is a duplicate of the previous element.
					// Why previous element l-1? Because we have already moved l one step to the right.
					if arr[l] == arr[l-1] {
						l++
					} else {
						break
					}
				}

			} else if s > ts {
				r--
			} else {
				l++
			}
		}
	}
	return triplets
}
