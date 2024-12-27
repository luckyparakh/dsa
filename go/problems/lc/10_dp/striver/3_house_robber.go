// https://leetcode.com/problems/house-robber/description/
package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := len(nums)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = -1
	}
	var solve func(n int) int
	solve = func(n int) int {
		if dp[n] != -1 {
			return dp[n]
		}
		max := 0
		for i := n + 2; i < l; i++ {
			temp := solve(i)
			if temp > max {
				max = temp
			}
		}
		dp[n] = max + nums[n]
		return dp[n]
	}

	first := dp[1]
	zero := solve(0)
	if dp[1] == -1 {
		first = solve(1)
	}
	if zero > first {
		return zero
	}
	return first
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1})) //2

}
