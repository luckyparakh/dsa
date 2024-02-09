// https://leetcode.com/problems/climbing-stairs/
// https://www.youtube.com/watch?v=mLfjzJsN8us&list=PLg0aancPZwRazLXPEW-vu517p3gXVCn0b&index=2
package main

import "fmt"

func climbStairs(n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	// fmt.Println(dp)
	var solve func(int) int
	solve = func(n int) int {
		// fmt.Println(n)
		if n == 0 || n == 1 {
			return 1
		}
		if n < 0 {
			return 0
		}
		prev1, prev2 := dp[n-1], dp[n-2]
		if prev1 == -1 {
			prev1 = solve(n - 1)
		}
		if prev2 == -1 {
			prev2 = solve(n - 2)
		}
		return prev1 + prev2
	}
	return solve(n)
}

func climbStairsBetter(n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			dp[i] = 1
		}
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairsBetter(3)) //3
	fmt.Println(climbStairsBetter(2)) //2
}
