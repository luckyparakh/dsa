// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
// https://leetcode.com/problems/split-array-largest-sum/submissions/

package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	minShipWeight := 1
	maxShipWeight := sum(weights)
	return func(s, e int) int {
		ans := 0
		for s <= e {
			shipWeight := s + (e-s)/2
			if isValid(weights, shipWeight, days) {
				ans = shipWeight
				e = shipWeight - 1
			} else {
				s = shipWeight + 1
			}
		}
		return ans
	}(minShipWeight, maxShipWeight)
}
func isValid(weights []int, shipWeight, days int) bool {
	s := 0
	daysNeeded := 0
	for _, weight := range weights {
		if weight > shipWeight {
			return false
		}
		s = s + weight
		if s > shipWeight {
			daysNeeded++
			s=weight
		}
	}
	// Edge case
	// At the end of loop, count of last weight is not calculated, for example s at end, will
	// less, equal or greater than shipWg.
	// If it is less or equal than shipWg then it will not enter 'if' at line 33 and count of that wg is missed
	// If it is greater than shipWg then it will enter 'if' at line 33 and count of till previous wgs are calculated
	// But still last wg is not contributing to count of days, hence below line of code
	if s > 0 {
		daysNeeded++
	}
	if daysNeeded > days {
		return false
	}
	return true
}
func sum(weights []int) int {
	s := 0
	for _, v := range weights {
		s = s + v
	}
	return s
}
func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
	fmt.Println(shipWithinDays([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(shipWithinDays([]int{5, 2, 1, 4, 6}, 3))
}
