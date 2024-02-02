// https://leetcode.com/problems/longest-palindromic-subsequence/description/
package main

import (
	"fmt"
	"math"
)

func longestPalindromeSubseq(s string) int {
	return lcsIn(s, reverse(s))
}
func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
func lcs(str1, str2 string) int {
	l1, l2 := len(str1)+1, len(str2)+1
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		arr := make([]int, l2)
		for j := 0; j < l2; j++ {
			arr[j] = -1
		}
		dp[i] = arr
	}

	var driver func(s1, s2 string, i1, i2 int) int
	driver = func(s1, s2 string, i1, i2 int) int {
		if i1 == 0 || i2 == 0 {
			return 0
		}
		if dp[i1][i2] != -1 {
			return dp[i1][i2]
		}
		if s1[i1-1] == s2[i2-1] {
			dp[i1][i2] = driver(s1, s2, i1-1, i2-1)
			return 1 + dp[i1][i2]
		} else {
			dp[i1][i2] = int(math.Max(float64(driver(s1, s2, i1, i2-1)),
				float64(driver(s1, s2, i1-1, i2))))
			return dp[i1][i2]
		}
	}
	return driver(str1, str2, len(str1), len(str2))
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

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab")) //4
	fmt.Println(longestPalindromeSubseq("cbbd"))  //2
}
