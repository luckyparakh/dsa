package main

import "fmt"

func searchNearlySortedArray(num []int, target int) int {
	l := len(num)
	s := 0
	e := l - 1
	for s <= e {
		m := s + (e-s)/2
		if m != 0 && num[m-1] == target {
			return m - 1
		}
		if m != l-1 && num[m+1] == target {
			return m + 1
		}
		if num[m] == target {
			return m
		}
		if num[m] > target {
			e = m - 2
		} else {
			s = m + 2
		}
	}
	return -1
}
func main() {
	fmt.Println(searchNearlySortedArray([]int{10, 3, 40, 20, 50, 80, 70}, 40))
	fmt.Println(searchNearlySortedArray([]int{10, 3, 40, 20, 50, 80, 70}, 0))
	fmt.Println(searchNearlySortedArray([]int{10, 3, 40, 20, 50, 80, 70}, 70))
	fmt.Println(searchNearlySortedArray([]int{10, 3, 40, 20, 50, 80, 70}, 10))
}
