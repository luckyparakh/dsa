// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/problem-challenge-1-quadruple-sum-to-target-medium
package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

func (s *Solution) searchQuadruplets(arr []int, target int) [][]int {
	var quadruplets [][]int
	sort.Ints(arr)
	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			l, r := j+1, len(arr)-1
			for l < r {
				sum := arr[i] + arr[j] + arr[l] + arr[r]
				// fmt.Println("Checking quadruplet:", arr[i], arr[j], arr[l], arr[r], "with sum:", sum)
				if sum == target {
					// fmt.Println("Found quadruplet:", arr[i], arr[j], arr[l], arr[r])
					quadruplets = append(quadruplets, []int{arr[i], arr[j], arr[l], arr[r]})
					r--
					l++
					for l < r && arr[r] == arr[r+1] {
						r--
					}
					for l < r && arr[l] == arr[l-1] {
						l++
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return quadruplets
}

func main() {
	s := Solution{}
	fmt.Println(s.searchQuadruplets([]int{0, 0, 0, 0, 0, 0}, 0))
	fmt.Println(s.searchQuadruplets([]int{4, 1, 2, -1, 1, -3}, 1))
}
