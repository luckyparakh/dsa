package main

import "fmt"

func smallestDivisor(nums []int, threshold int) int {
	min := -1
	s, e := 1, findMax(nums)

	for s <= e {
		m := s + (e-s)/2
		if isValid(nums, m, threshold) {
			min = m
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return min
}
func isValid(nums []int, divisor, threshold int) bool {
	// fmt.Println(divisor)
	sum := 0
	for _, num := range nums {
		r := num / divisor
		if num%divisor == 0 {
			sum += r
		} else {
			sum += r + 1
		}
	}
	return sum <= threshold
}
func findMax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func main() {
	fmt.Println("ans ",smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println("ans ",smallestDivisor([]int{44, 22, 33, 11, 1}, 5))
}
