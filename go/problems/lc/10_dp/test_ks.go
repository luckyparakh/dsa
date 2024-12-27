package main

func knapsackRecursive(wt []int, vs []int, capacity int, n int) int {
	if n == 0 || capacity == 0 {
		return 0
	}

	if wt[n-1] > capacity {
		return knapsackRecursive(wt, vs, capacity, n-1)
	}

	return max(vs[n-1]+knapsackRecursive(wt, vs, capacity-wt[n-1], n-1),
		knapsackRecursive(wt, vs, capacity, n-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
