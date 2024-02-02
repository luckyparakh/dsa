// https://practice.geeksforgeeks.org/problems/rod-cutting0840/1
// https://www.youtube.com/watch?v=SZqAQLjDsag&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=14

package main

import (
	"fmt"
	"math"
)
/*
Unbounded Knapsack

The cutRodRecur function is a recursive implementation of the rod cutting problem. 
It calculates the maximum obtainable value by cutting a rod of length n into smaller pieces, 
given the prices of different lengths of the rod.

Inputs
price: an array of integers representing the prices of different lengths of the rod
n: an integer representing the length of the rod

Flow
The function defines a nested function called driver that takes three parameters: price, size, and element.
If the size of the rod is 0 or the element is 0, it means there are no more cuts to be made, so the function returns 0.
If the element is less than or equal to the size, it means the current length of the rod can be cut, 
so the function calculates the maximum value by either taking the current length or not taking it.
The function recursively calls itself to calculate the maximum value for the remaining length of the rod 
and the remaining elements.
The function returns the maximum value obtained by either taking the current length or not taking it.
*/
func cutRodRecur(price []int, n int) int {
	var driver func([]int, int, int) int
	driver = func(price []int, size, element int) int {
		if size == 0 {
			return 0
		}
		if element == 0 {
			return 0
		}
		if element <= size {
			taken := price[element-1] + driver(price, size-element, element)
			notTaken := driver(price, size, element-1)
			return int(math.Max(float64(taken), float64(notTaken)))
		}
		return driver(price, size, element-1)

	}
	return driver(price, n, n)
}

func cutRodRecurM(price []int, n int) int {
	size := n
	lenPrice := len(price)
	dp := make([][]int, lenPrice+1) // Create a 2D slice for dynamic programming
	for i := 0; i < lenPrice+1; i++ {
		s := make([]int, size+1) // Create a slice for each row of the 2D slice
		for j := 0; j < size+1; j++ {
			s[j] = -1
		}
		dp[i] = s
	}
	var driver func([]int, int, int) int
	driver = func(price []int, size, element int) int {
		if size == 0 {
			dp[element][size] = 0
			return 0
		}
		if element == 0 {
			dp[element][size] = 0
			return 0
		}
		if dp[element][size] != -1 {
			return dp[element][size]
		}
		if element <= size {
			taken := price[element-1] + driver(price, size-element, element)
			notTaken := driver(price, size, element-1)
			dp[element][size] = int(math.Max(float64(taken), float64(notTaken)))
			return dp[element][size]
		}
		dp[element][size] = driver(price, size, element-1)
		return dp[element][size]
	}
	return driver(price, n, n)
}
func cutRod(price []int, n int) int {
	size := n
	lenPrice := len(price)
	dp := make([][]int, lenPrice+1) // Create a 2D slice for dynamic programming
	for i := 0; i < lenPrice+1; i++ {
		s := make([]int, size+1) // Create a slice for each row of the 2D slice
		dp[i] = s
	}
	for i := 0; i < lenPrice+1; i++ {
		for j := 0; j < size+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				if i <= j {
					dp[i][j] = int(math.Max(float64(price[i-1]+dp[i][j-i]),
						float64(dp[i-1][j])))
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[lenPrice][n]
}

func main() {
	fmt.Println(cutRod([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8))
	fmt.Println(cutRodRecur([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8))
	fmt.Println(cutRodRecurM([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8))
}
