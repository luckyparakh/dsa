package main

import (
	"fmt"
	"strconv"
)

func findFloorCeil(num []int, t int) []int {
	s := 0
	e := len(num) - 1
	op := []int{}
	for s <= e {
		m := s + (e-s)/2
		if m == 0 && num[m] == t {
			op = append(op, -1, m+1)
			break
		}
		if m == len(num)-1 && num[m] == t {
			op = append(op, m-1, -1)
			break
		}
		if num[m] == t {
			op = append(op, m-1, m+1)
			break
		}
		if num[m] > t {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	if len(op) == 0 {
		if s == len(num) {
			op = append(op, e, -1)
		} else if e == -1 {
			op = append(op, -1, s)
		} else {
			op = append(op, e, s)
		}

	}
	return op
}
func main() {
	fmt.Println(findFloorCeil([]int{1, 3, 5, 6}, 5))
	fmt.Println(findFloorCeil([]int{1, 3, 5, 6}, 7))
	fmt.Println(findFloorCeil([]int{1, 3, 5, 6}, 2))
	fmt.Println(findFloorCeil([]int{1, 3, 5, 6}, 1))
	fmt.Println(findFloorCeil([]int{1, 3, 5, 6}, 6))
}
