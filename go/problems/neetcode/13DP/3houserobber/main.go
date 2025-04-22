package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return helper(nums, memo, 0)
}

func helper(nums, memo []int, n int) int {
	if n >= len(nums) {
		return 0
	}
	if memo[n] != -1 {
		return memo[n]
	}
	memo[n] = max(nums[n]+helper(nums, memo, n+2), helper(nums, memo, n+1))
	return memo[n]
}
