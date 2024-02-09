// https://leetcode.com/problems/min-cost-climbing-stairs/

package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	size := 1001
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = -1
	}
	var solve func(index int) int
	solve = func(index int) int {
		if index >= len(cost) {
			return 0
		}
		if dp[index] != -1 {
			return dp[index]
		}
		l, r := dp[index+1], dp[index+2]
		if l == -1 {
			l = solve(index + 1)
		}
		if r == -1 {
			r = solve(index + 2)
		}
		left := l + cost[index]
		right := r + cost[index]
		dp[index] = int(math.Min(float64(left), float64(right)))
		return dp[index]
	}

	return int(math.Min(float64(solve(0)), float64(solve(1))))
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
