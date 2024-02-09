// https://www.geeksforgeeks.org/problems/palindromic-patitioning4845/1

package main

import (
	"fmt"
	"math"
)

func palindromicPartition(str string) int {
	size := 501
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		t := make([]int, size)
		for j := 0; j < size; j++ {
			t[j] = -1
		}
		dp[i] = t
	}
	var solve func(string, int, int) int
	solve = func(s string, i, j int) int {
		if i == j || isPalin(s[i:j+1]) {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		min := math.MaxInt
		for k := i; k < j; k++ {
			left, right := dp[i][k], dp[k+1][j]
			if dp[i][k] == -1 {
				left = solve(s, i, k)
			}
			if dp[k+1][j] == -1 {
				right = solve(s, k+1, j)
			}
			tempAns := left + right + 1
			if min > tempAns {
				min = tempAns
			}
		}
		dp[i][j] = min
		return dp[i][j]
	}
	return solve(str, 0, len(str)-1)
}

func isPalin(s string) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {

			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(palindromicPartition("ababbbabbababa")) //3
	fmt.Println(palindromicPartition("aaabba"))         //1
}
