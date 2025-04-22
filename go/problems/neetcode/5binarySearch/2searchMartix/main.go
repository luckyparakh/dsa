package main

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	c := len(matrix[0])
	s := 0
	e := r * c
	for s <= e {
		m := s + (e-s)/2
		v := matrix[m/c][m%c]
		if target == v {
			return true
		} else if target > v {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return false
}
