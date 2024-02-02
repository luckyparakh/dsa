// https://leetcode.com/problems/partition-equal-subset-sum/

package main

import "fmt"

/*
The code calculates the sum of all the numbers in the nums slice.
If the sum is even (sum%2 == 0), it calls the isSubsetSumDP function with the length of nums, nums itself, and half of the sum.
Otherwise, it returns false.
*/
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 0 {
		return isSubsetSumDP(len(nums), nums, sum/2)
	}
	return false
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

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))    //true
	fmt.Println(canPartition([]int{2, 4, 6, 8, 10})) //false
}
