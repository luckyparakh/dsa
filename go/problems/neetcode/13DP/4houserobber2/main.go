package main

/*
	Create Two Linear Subarrays:

Subarray 1 (Robbing the first house, excluding the last):
Create a subarray that includes nums[0] to nums[len(nums)-2]. This represents the scenario where you rob the first house, so you cannot rob the last.
Subarray 2 (Robbing the last house, excluding the first):
Create a subarray that includes nums[1] to nums[len(nums)-1]. This represents the scenario where you rob the last house, so you cannot rob the first.
*/
func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	nums1 := nums[:l-1]
	nums2 := nums[1:]
	memo1 := make([]int, len(nums1))
	for i := range memo1 {
		memo1[i] = -1
	}
	memo2 := make([]int, len(nums2))
	for i := range memo2 {
		memo2[i] = -1
	}
	return max(helper(nums1, memo1, 0), helper(nums2, memo2, 0))
	// return max(helper1(nums1), helper1(nums2))
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

// func helper1(nums []int) int {
// 	l := len(nums)
// 	dp := make([]int, l)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = -1
// 	}
// 	if l == 1 {
// 		dp[0] = nums[0]
// 	}
// 	if l == 2 {
// 		dp[1] = max(nums[0], nums[1])
// 	}

// 	for i := 2; i < l; i++ {
// 		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
// 	}
// 	return dp[l]
// }
