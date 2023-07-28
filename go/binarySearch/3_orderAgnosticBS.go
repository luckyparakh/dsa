package main

import "fmt"

func binarySearchR(n []int, num int) int {
	fmt.Println("BSR")
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
func binarySearch(n []int, num int) int {
	fmt.Println("BS")
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
func findOrder(n []int, num int) int {
	s := 0
	e := len(n) - 1
	if n[s] <= n[e] {
		return binarySearch(n, num)
	}
	return binarySearchR(n, num)
}
func main() {
	fmt.Println(findOrder([]int{1, 2, 3, 4, 5, 6}, 4))
	fmt.Println(findOrder([]int{1, 2, 3, 4, 5, 6}, 10))
	fmt.Println(findOrder([]int{6, 5, 4, 3, 2, 1}, 5))
	fmt.Println(findOrder([]int{6, 5, 4, 3, 2, 1}, 10))
	fmt.Println(findOrder([]int{7,7,7}, 7))
	fmt.Println(findOrder([]int{7,7,7}, 10))
}
