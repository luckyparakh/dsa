// https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/find-nonduplicate-number-instances-easy

package main

type Solution struct {
}

func (sol *Solution) moveElements(arr []int) int {
	insertIndex, scanIndex := 0, 0
	// Why not left < len(arr)? Because right is the pointer that goes through the array.
	for scanIndex < len(arr) {
		if arr[insertIndex] == arr[scanIndex] {
			scanIndex++
		} else {
			insertIndex++
			arr[insertIndex] = arr[scanIndex]
		}
	}
	return insertIndex + 1
}
