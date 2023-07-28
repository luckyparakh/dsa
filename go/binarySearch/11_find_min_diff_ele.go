package main

// Find the minimum difference element in a sorted array

import "fmt"

func findMinDiff(a []int, t int) int {
	l := len(a)
	if t < a[0] {
		return a[0]
	}
	if t > a[l-1] {
		return a[l-1]
	}
	s, e := 0, l-1

	for s <= e {
		m := s + (e-s)/2
		if a[m] == t {
			return a[m]
		}
		if a[m] > t {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	// e point to lower value and s points to higher value
	if t-a[e] > a[s]-t {
		return a[s]
	}
	return a[e]
}
func main() {
	fmt.Println(findMinDiff([]int{2, 5, 10, 12, 15}, 6))
	fmt.Println(findMinDiff([]int{2, 5, 10, 12, 15}, 10))
	fmt.Println(findMinDiff([]int{2, 5, 10, 12, 15}, 1))
	fmt.Println(findMinDiff([]int{2, 5, 10, 12, 15}, 16))
}
