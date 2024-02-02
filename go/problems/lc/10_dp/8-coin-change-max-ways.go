// https://www.youtube.com/watch?v=I4UR2T6Ro3w&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=15
// https://practice.geeksforgeeks.org/problems/coin-change2448/1
package main

import "fmt"

func count(coins []int, N, sum int) int64 {
	dp := make([][]int64, N+1) // Create a 2D slice for dynamic programming
	for i := 0; i < N+1; i++ {
		s := make([]int64, sum+1) // Create a slice for each row of the 2D slice
		dp[i] = s
	}
	for i := 0; i < N+1; i++ {
		for j := 0; j < sum+1; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else {
				if coins[i-1] <= j {
					dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	fmt.Print(dp)
	return dp[N][sum]
}

func main() {
	fmt.Println(count([]int{1, 2, 3}, 3, 4))
}
