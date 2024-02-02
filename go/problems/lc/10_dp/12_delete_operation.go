// https://leetcode.com/problems/delete-operation-for-two-strings/description/
// https://www.youtube.com/watch?v=CFwCCNbRuLY&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=27
package main

import "math"

func minDistance(word1 string, word2 string) int {
	lcs := lcsIn(word1, word2)
	return len(word1) - lcs + len(word2) - lcs
}

func lcsIn(str1, str2 string) int {
	l1, l2 := len(str1)+1, len(str2)+1
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		arr := make([]int, l2)
		for j := 0; j < l2; j++ {
			arr[j] = -1
		}
		dp[i] = arr
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j])))
			}
		}
	}
	return dp[l1-1][l2-1]
}
