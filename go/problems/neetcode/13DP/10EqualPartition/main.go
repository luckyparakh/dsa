package main

func canPartition(nums []int) bool {
	l := len(nums) - 1
	sum := 0
	for i := 0; i <= l; i++ {
		sum += nums[i]
	}
	if sum%2 == 0 {
		return subsetSum(nums, sum/2)
	}
	return false
}
func subsetSum(arr []int, sum int) bool {
	l := len(arr)
	dp := make([][]bool, l+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i < l+1; i++ {
		for s := 0; s < sum+1; s++ {
			if s == 0 {
				dp[i][s] = true
			} else if i == 0 {
				dp[i][s] = false
			} else {
				if arr[i-1] <= s {
					dp[i][s] = dp[i-1][s-arr[i-1]] || dp[i-1][s]
				} else {
					dp[i][s] = dp[i-1][s]
				}
			}
		}
	}
	// fmt.Println(dp)
	return dp[l][sum]
}

func minSubSetDiff(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ans := 1000000
	lastRow := findSubset(arr, sum)
	for i := 0; i <= len(lastRow)/2; i++ {
		if lastRow[i] {
			val := sum - 2*i
			ans = min(ans, val)
		}
	}
	return ans
}
func findSubset(arr []int, sum int) []bool {
	l := len(arr)
	dp := make([][]bool, l+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	for n := 0; n < l+1; n++ {
		for s := 0; s < sum+1; s++ {
			if s == 0 {
				dp[n][s] = true
			} else if n == 0 {
				dp[n][s] = false
			} else {
				if arr[n-1] <= s {
					dp[n][s] = dp[n-1][s-arr[n-1]] || dp[n-1][s]
				} else {
					dp[n][s] = dp[n-1][s]
				}
			}
		}
	}
	return dp[l]
}
