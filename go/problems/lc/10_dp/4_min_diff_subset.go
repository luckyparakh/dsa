// https://practice.geeksforgeeks.org/problems/minimum-sum-partition3317/1
// https://www.youtube.com/watch?v=-GtpxG6l_Mc&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=10
package main

import "fmt"
/*
s1+s2=sum
s1-s2=diff (we have find min diff)
With these 2 eq: diff=sum-2s2 
And to find min diff, s2 should be max but not be greater than half of Sum
*/
func minDifference(arr []int, n int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	boolArr := isSubsetSumDP(len(arr), arr, sum)
	// fmt.Println(boolArr)
	largestTrueIndex := 0
	for i := 1; i <= (len(boolArr)-1)/2; i++ {

		if boolArr[i] {
			if i > largestTrueIndex {
				largestTrueIndex = i
			}
		}
	}
	// fmt.Println(largestTrueIndex)
	return sum - 2*largestTrueIndex
}

func isSubsetSumDP(N int, arr []int, sum int) []bool {
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
	// Why returning last row? Because last row means with N elements, how many subset sum can be created.
	return dp[N] // Return the result stored in dp[N]
}

func main() {
	fmt.Println(minDifference([]int{1, 6, 11, 5}, 4))            // 1
	fmt.Println(minDifference([]int{1, 4}, 2))                   // 3
	fmt.Println(minDifference([]int{5, 6, 6, 5, 7, 6, 4, 7}, 8)) // 0
}
