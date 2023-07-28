// Find 1st occurrence of 1 in infinite array
package main

import "fmt"

func searchInfiniteArray(a []int, t int) int {
	s := 0
	e := 1
	for {
		if a[e] >= t {
			break
		}
		s = e
		e = e * 2
	}
	return bs(a, s, e, t)
}
func bs(a []int, s, e, t int) int {
	for s <= e {
		m := s + (e-s)/2
		if a[m] == t {
			e = m - 1
		}
		if a[m] > t {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return e + 1
}
func main() {
	fmt.Println(searchInfiniteArray([]int{0, 0, 0, 0, 0, 0, 1, 1, 1}, 1))
	fmt.Println(searchInfiniteArray([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}, 1))
	fmt.Println(searchInfiniteArray([]int{1, 1, 1}, 1))
}
