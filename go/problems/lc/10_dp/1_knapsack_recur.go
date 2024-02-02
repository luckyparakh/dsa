// https://practice.geeksforgeeks.org/problems/0-1-knapsack-problem0945/1
// https://www.youtube.com/watch?v=kvyShbFVaY8&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=3
// https://www.youtube.com/watch?v=fJbIuhs24zQ&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=4
// https://www.youtube.com/watch?v=ntCGbPMeqgg&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=5

package main

import "math"

// knapSack is a function that solves the 0/1 Knapsack problem using recursion.
//
// Inputs:
// - W: an integer representing the maximum weight capacity of the knapsack.
// - wt: a slice of integers representing the weights of the items.
// - val: a slice of integers representing the values of the items.
// - n: an integer representing the number of items.
//
// Returns:
// An integer representing the maximum value that can be obtained by selecting items to fit within the knapsack's weight capacity.

func knapSack(W int, wt []int, val []int, n int) int {
	if n == 0 || W == 0 {
		return 0
	}
	if wt[n-1] <= W {
		return int(math.Max(float64(val[n-1]+knapSack(W-wt[n-1], wt, val, n-1)),
			float64(knapSack(W, wt, val, n-1))))
	}
	return knapSack(W, wt, val, n-1)
}

// Optimized code with Memoization
func knapSack1(W int, wt []int, val []int, n int) int {
	// Create slice of n+1 and W+1 size, as 0th row is used for initialization
	dp := make([][]int, n+1) // Create a 2D slice to store the memoization table
	for i := 0; i < n+1; i++ {
		w := make([]int, W+1) // Create a 1D slice to store the values for each weight
		for j := 0; j < W+1; j++ {
			w[j] = -1 // Initialize the values in the slice to -1
		}
		dp[i] = w // Assign the 1D slice to the corresponding row in the 2D slice
	}
	var driver func(int, []int, []int, int) int // Declare a recursive function called "driver"
	driver = func(W int, wt []int, val []int, n int) int {
		if n == 0 || W == 0 {
			dp[n][W] = 0 // Base case: if there are no items or the weight capacity is 0, set the value in the memoization table to 0
			return dp[n][W]
		}
		if dp[n][W] != -1 {
			return dp[n][W] // If the value is already computed and stored in the memoization table, return it
		}
		if wt[n-1] <= W {
			dp[n][W] = int(math.Max(float64(val[n-1]+driver(W-wt[n-1], wt, val, n-1)),
				float64(driver(W, wt, val, n-1)))) // Recursive case: compute the maximum value by either including or excluding the current item
			return dp[n][W]
		}
		return driver(W, wt, val, n-1) // Recursive case: exclude the current item and move to the next item
	}
	return driver(W, wt, val, n) // Call the "driver" function to solve the knapsack problem using memoization
}

func knapSackIter(W int, wt []int, val []int, n int) int {
	// Create slice of n+1 and W+1 size, as 0th row is used for initialization
	dp := make([][]int, n+1) // Create a 2D slice to store the memoization table
	for i := 0; i < n+1; i++ {
		w := make([]int, W+1) // Create a 1D slice to store the values for each weight
		for j := 0; j < W+1; j++ {
			w[j] = -1 // Initialize the values in the slice to -1
		}
		dp[i] = w // Assign the 1D slice to the corresponding row in the 2D slice
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			// Initialization same as base condition in recursion
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				if wt[i-1] <= j {
					// If the weight of the current item is less than or equal to the remaining weight capacity,
					// we have two choices:
					// 1. Include the current item and add its value to the value obtained by selecting items with the remaining weight capacity.
					// 2. Exclude the current item and consider the value obtained by selecting items with the same weight capacity.
					// We choose the maximum value among these two choices.
					dp[i][j] = int(math.Max(float64(val[i-1]+dp[i-1][j-wt[i-1]]),
						float64(dp[i-1][j])))
				} else {
					// If the weight of the current item is greater than the remaining weight capacity,
					// we cannot include the current item. So, we consider the value obtained by selecting items with the same weight capacity.
					dp[i][j] = dp[i-1][j]
				}
			}

		}
	}
	return dp[n][W]
}
