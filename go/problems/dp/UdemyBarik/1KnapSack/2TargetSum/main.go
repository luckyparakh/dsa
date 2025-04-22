package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	return helper(nums, target, 0, 0)
}

func helper(nums []int, target, s, i int) int {
	// fmt.Println(i,target)
	if target == s && i == len(nums) {
		return 1
	}
	if i == len(nums) {
		return 0
	}
	return helper(nums, target, s+nums[i], i+1) + helper(nums, target, s-nums[i], i+1)
}

func findTargetSumWaysDp(nums []int, target int) int {
	l := len(nums)
	total := 0
	for i := range nums {
		total += nums[i]
	}
	dp := make([][]int, l+1)
	for i := range dp {
		t := make([]int, (2*total)+1)
		for j := range t {
			t[j] = -1
		}
		dp[i] = t
	}
	fmt.Println(dp)

	return helper1(nums, target,total, 0, 0, dp)
}

func helper1(nums []int, target,total, s, i int, dp [][]int) int {
	l := len(nums)
	if target == s && i == l {
		return 1
	}
	if i == l {
		return 0
	}
	if dp[i][s+total] != -1 {
		return dp[i][s+total]
	}
	dp[i][s+total]= helper1(nums, target,total, s+nums[i], i+1, dp) + helper1(nums, target,total, s-nums[i], i+1, dp)
	return dp[i][s+total]
}
