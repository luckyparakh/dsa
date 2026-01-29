// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/squaring-a-sorted-array-easy

package main

type Solution struct{}

func (s *Solution) makeSquares(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)
	left, right := 0, n-1
	placement := n - 1
	// Why left <= right? Because we need to process all elements, including when left == right.
	for left <= right {
		rightSq := arr[right] * arr[right]
		leftSq := arr[left] * arr[left]

		if leftSq < rightSq {
			squares[placement] = rightSq
			right--
		} else {
			squares[placement] = leftSq
			left++
		}
		placement--
	}
	return squares
}
