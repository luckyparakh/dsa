package main

import "fmt"

func binarySearch(n []int, num int) int {
	s := 0
	e := len(n) - 1
	for s <= e {
		mid := s + (e-s)/2
		if n[mid] == num {
			return mid
		}
		if n[mid] > num {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return -1
}
func main() {
	fmt.Println(binarySearch([]int{6, 5, 4, 3, 2, 1}, 5))
	fmt.Println(binarySearch([]int{6, 5, 4, 3, 2, 1}, 1))
	fmt.Println(binarySearch([]int{6, 5, 4, 3, 2, 1}, 10))
}
