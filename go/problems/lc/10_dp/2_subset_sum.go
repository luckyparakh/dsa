// https://practice.geeksforgeeks.org/problems/subset-sum-problem-1611555638/1
// https://www.youtube.com/watch?v=_gPcYovP7wc&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=7

package main

func isSubsetSum(N int, arr []int, sum int) bool {
	if sum == 0 { // If the sum is zero, return true
		return true
	}
	if N == 0 { // If N is zero, return false
		return false
	}

	if arr[N-1] <= sum { // If the last element of arr is less than or equal to sum
		return isSubsetSum(N-1, arr, sum-arr[N-1]) || isSubsetSum(N-1, arr, sum) // Recursively check if the sum can be achieved by including or excluding the last element
	}
	return isSubsetSum(N-1, arr, sum) // Recursively check if the sum can be achieved by excluding the last element
}

func isSubsetSumDP(N int, arr []int, sum int) bool {
	dp := make([][]bool, N+1) // Create a 2D slice for dynamic programming
	for i := 0; i < N+1; i++ {
		s := make([]bool, sum+1) // Create a slice for each row of the 2D slice
		dp[i] = s
	}
	for i := 0; i < N+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 || j == 0 {
				// Although it is not needed but wrote for clarity
				if i == 0 { // If i is zero, set dp[i][j] to false
					dp[i][j] = false
				}
				if j == 0 { // If j is zero, set dp[i][j] to true
					dp[i][j] = true
				}
			} else {
				if arr[i-1] <= j { // If the current element can be included and the remaining sum is valid
					dp[i][j] = dp[i-1][j-arr[i-1]] || dp[i-1][j] // Update dp[i][j] based on whether including or excluding the current element leads to a valid sum
				} else {
					dp[i][j] = dp[i-1][j] // If the current element cannot be included, update dp[i][j] based on the previous row
				}
			}
		}
	}
	return dp[N][sum] // Return the result stored in dp[N][sum]
}
