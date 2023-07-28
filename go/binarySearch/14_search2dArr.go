// https://leetcode.com/problems/search-a-2d-matrix/

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rl := len(matrix[0])
	l := len(matrix) * rl
	s := 0
	e := l - 1
	for s <= e {
		mid := s + (e-s)/2
		i, j := mid/rl, mid%rl
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return false
}
func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}, 3))
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}, 35))
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7, 23, 30, 34, 60},
	}, 34))

}
