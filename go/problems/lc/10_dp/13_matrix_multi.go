// https://www.geeksforgeeks.org/problems/matrix-chain-multiplication0303/1
package main

import "math"

func matrixMultiplication(n int, arr []int) int {
	size := 101
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		t := make([]int, size)
		for j := 0; j < size; j++ {
			t[j] = -1
		}
		dp = append(dp, t)
	}
	var solve func([]int, int, int) int
	solve = func(arr []int, i, j int) int {
		if i >= j {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		min := math.MaxInt
		for k := i; k < j; k++ {
			left, right := dp[i][k], dp[k+1][j]
			if dp[i][k] == -1 {
				left = solve(arr, i, k)
			}
			if dp[k+1][j] == -1 {
				right = solve(arr, k+1, j)
			}
			temp := left + right + (arr[i-1] * arr[k] * arr[j])
			if min > temp {
				min = temp
			}
		}
		dp[i][j] = min
		return dp[i][j]
	}
	return solve(arr, 1, n-1)
}
