package main

import "fmt"

func sortArr(arr []int) []int {
	fmt.Println(arr)
	if len(arr) == 1 {
		return arr
	}
	return insert(sortArr(arr[:len(arr)-1]), arr[len(arr)-1])
}
func insert(arr []int, ele int) []int {
	// fmt.Println(arr,ele)
	arrI := []int{}
	inserted := false
	for _, v := range arr {
		if ele < v && !inserted {
			arrI = append(arrI, ele)
			inserted = true
		}
		arrI = append(arrI, v)
	}
	if !inserted {
		arrI = append(arrI, ele)
	}
	// fmt.Println(arrI)
	return arrI
}
func insertBS(arr []int, ele int) []int {
	// fmt.Println(arr, ele)
	s := 0
	e := len(arr) - 1
	
	for s <= e {
		m := s + (e-s)/2
		if arr[m] > ele {
			e = m - 1
		} else if ele > arr[m] {
			s = m + 1
		} else {
			break
		}
		// fmt.Println(s, e, m)
	}
	position := e + 1
	// arrI := make([]int, len(arr)+1)
	fmt.Println(position)
	arrI := []int{}
	if position == len(arr) {
		arrI = arr
		arrI = append(arrI, ele)
	} else {
		for k, v := range arr {
			if position == k {
				arrI = append(arrI, ele)
			}
			arrI = append(arrI, v)
		}
	}
	// fmt.Println(arrI)
	return arrI
}
func main() {
	// fmt.Println(insert([]int{1}, 2))
	fmt.Println(insertBS([]int{0, 15, 17, 20, 25, 27}, 16))
	fmt.Println(insertBS([]int{8, 15, 17, 20, 25, 27}, 6))
	fmt.Println(insertBS([]int{8, 17}, 15))
	fmt.Println(insertBS([]int{8}, 17))
	fmt.Println(sortArr([]int{8, 17, 15, 27, 25, 20}))
}
