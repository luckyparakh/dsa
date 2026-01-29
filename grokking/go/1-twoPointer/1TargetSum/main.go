package main

type Solution struct {
}

// search - finds a pair in arr with a given targetSum.
func (sol *Solution) search(arr []int, targetSum int) [2]int {
	left, right := 0, len(arr)-1
	// Why not left <= right? Because when left == right, we are looking at the same element,
	// and we need a pair of different elements to satisfy the condition.
	for left < right {
		sum := arr[left] + arr[right]
		if sum == targetSum {
			return [2]int{left, right}
		} else if sum > targetSum {
			right--
		} else {
			left++
		}

	}

	return [2]int{-1, -1} // pair not found
}
