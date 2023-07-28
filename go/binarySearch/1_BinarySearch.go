// https://www.youtube.com/watch?v=j7NodO9HIbk&list=PL_z_8CaSLPWeYfhtuKHj-9MpYb6XQJ_f2&index=1
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
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 4))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 10))
}
