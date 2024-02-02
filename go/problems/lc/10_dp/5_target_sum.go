package main

// https://leetcode.com/problems/target-sum/description/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	/*
		s1+s2=sum
		s1-s2=diff
		2s2=sum-diff
	*/
	if target > sum && (sum-target)%2 != 0 {
		return 0
	}
	return isSubsetSumDP(len(nums), nums, (sum-target)/2)
}

func isSubsetSumDP(N int, arr []int, sum int) int {
	dp := make([][]int, N+1) // Create a 2D slice for dynamic programming
	for i := 0; i < N+1; i++ {
		s := make([]int, sum+1) // Create a slice for each row of the 2D slice
		dp[i] = s
	}
	for i := 0; i < N+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 || j == 0 {
				// Although it is not needed but wrote for clarity
				if i == 0 { // If i is zero, set dp[i][j] to 0
					dp[i][j] = 0
				}
				if j == 0 { // If j is zero, set dp[i][j] to 1
					dp[i][j] = 1
				}
			} else {
				if arr[i-1] <= j { // If the current element can be included and the remaining sum is valid
					dp[i][j] = dp[i-1][j-arr[i-1]] + dp[i-1][j] // Update dp[i][j] based on whether including or excluding the current element leads to a valid sum
				} else {
					dp[i][j] = dp[i-1][j] // If the current element cannot be included, update dp[i][j] based on the previous row
				}
			}
		}
	}
	return dp[N][sum] // Return the result stored in dp[N][sum]
}
