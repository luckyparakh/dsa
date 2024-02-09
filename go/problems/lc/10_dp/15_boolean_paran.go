// https://www.geeksforgeeks.org/problems/boolean-parenthesization5610/1?itm_source=geeksforgeeks&itm_medium=article&itm_campaign=bottom_sticky_on_article
// https://www.youtube.com/watch?v=pGVguAcWX4g&t=1108s

package main

import (
	"fmt"
	"strconv"
)

/*
countWays is a function that calculates the number of ways to parenthesize a boolean expression to obtain a desired result.

Parameters:
- n (int): The length of the boolean expression.
- s (string): The boolean expression.

Returns:
- int: The number of ways to parenthesize the boolean expression.

Description:
The countWays function uses dynamic programming to solve the problem. It maintains a map (mp) to store previously calculated results for subproblems. It also defines a recursive helper function (solve) that takes the boolean expression, the starting and ending indices, and a boolean flag indicating whether the desired result is true or false. The solve function recursively calculates the number of ways to parenthesize the subexpression from index i to j to obtain the desired result.

The solve function uses the following logic:
- If the starting index is greater than the ending index, there are no ways to parenthesize the subexpression, so it returns 0.
- If the starting index is equal to the ending index, there is only one character in the subexpression. If the desired result is true and the character is 'T', or if the desired result is false and the character is 'F', there is one way to parenthesize the subexpression to obtain the desired result. Otherwise, there are no ways, so it returns 0.
- The function checks if the result for the current subproblem has already been calculated and stored in the map. If so, it returns the stored result.
- The function initializes a variable (ans) to store the number of ways to parenthesize the subexpression.
- It iterates over all possible operators in the subexpression (every second character starting from the second index).
- For each operator, it recursively calculates the number of ways to parenthesize the left and right subexpressions, both for the true and false results.
- It applies the operator to the left and right results based on the current operator.
- It updates the answer by adding the calculated values.
- Finally, it stores the calculated result in the map and returns the result modulo 1003.

The countWays function calls the solve function with the initial parameters (the entire boolean expression, starting index 0, ending index len(s)-1, and desired result true) and returns the result.

Example:
countWays(7, "T|T&F^T") returns 4
countWays(5, "T^F|F") returns 2
*/

func countWays(n int, s string) int {
	mp := map[string]int{}
	var solve func(string, int, int, bool) int

	solve = func(s string, i, j int, isTrue bool) int {
		if i > j {
			return 0
		}
		if i == j {
			if isTrue && s[i] == 'T' || !isTrue && s[i] == 'F' {
				return 1
			} else {
				return 0
			}
		}
		key := getKey(i, j, isTrue)
		if val, found := mp[key]; found {
			return val
		}
		ans := 0
		for k := i + 1; k < j; k = k + 2 {
			lt, lf, rt, rf := -1, -1, -1, -1
			if val, found := mp[getKey(i, k-1, true)]; found {
				lt = val
			} else {
				lt = solve(s, i, k-1, true)
			}
			if val, found := mp[getKey(i, k-1, false)]; found {
				lf = val
			} else {
				lf = solve(s, i, k-1, false)
			}
			if val, found := mp[getKey(k+1, j, true)]; found {
				rt = val
			} else {
				rt = solve(s, k+1, j, true)
			}
			if val, found := mp[getKey(k+1, j, false)]; found {
				rf = val
			} else {
				rf = solve(s, k+1, j, false)
			}

			if s[k] == '&' {
				if isTrue {
					ans = ans + lt*rt
				} else {
					ans = ans + lt*rf + lf*rf + lf*rt
				}
			}
			if s[k] == '^' {
				if isTrue {
					ans = ans + lt*rf + lf*rt
				} else {
					ans = ans + lf*rf + lt*rt
				}
			}
			if s[k] == '|' {
				if isTrue {
					ans = ans + lt*rt + lt*rf + lf*rt
				} else {
					ans = ans + lf*rf
				}
			}
		}
		mp[key] = ans % 1003
		return ans % 1003
	}
	return solve(s, 0, len(s)-1, true)
}
func countWays1(n int, s string) int {
	/*
	Using a three-dimensional slice can improve the performance and memory efficiency of the code. 
	Maps have an overhead in terms of memory usage and lookup time, 
	while slices provide direct access to elements and can be more efficient for this use case.
	*/
	const modulo = 1003
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -1
			dp[i][j][1] = -1
		}
	}
	var solve func(string, int, int, bool) int

	solve = func(s string, i, j int, isTrue bool) int {
		if i > j {
			return 0
		}
		if i == j {
			if isTrue && s[i] == 'T' || !isTrue && s[i] == 'F' {
				return 1
			} else {
				return 0
			}
		}

		if isTrue {
			if dp[i][j][0] != -1 {
				return dp[i][j][0]
			}
		} else {
			if dp[i][j][1] != -1 {
				return dp[i][j][1]
			}
		}
		ans := 0
		for k := i + 1; k < j; k = k + 2 {
			lt, lf, rt, rf := dp[i][k-1][0], dp[i][k-1][1], dp[k+1][j][0], dp[k+1][j][1]
			if dp[i][k-1][0] != -1 {
				lt = dp[i][k-1][0]
			} else {
				lt = solve(s, i, k-1, true)
			}
			if dp[i][k-1][1] != -1 {
				lf = dp[i][k-1][1]
			} else {
				lf = solve(s, i, k-1, false)
			}
			if dp[k+1][j][0] != -1 {
				rt = dp[k+1][j][0]
			} else {
				rt = solve(s, k+1, j, true)
			}
			if dp[k+1][j][1] != -1 {
				rf = dp[k+1][j][1]
			} else {
				rf = solve(s, k+1, j, false)
			}
			switch s[k] {
			case '&':
				if isTrue {
					ans = ans + lt*rt
				} else {
					ans = ans + lt*rf + lf*rf + lf*rt
				}
			case '^':
				if isTrue {
					ans = ans + lt*rf + lf*rt
				} else {
					ans = ans + lf*rf + lt*rt
				}
			case '|':
				if isTrue {
					ans = ans + lt*rt + lt*rf + lf*rt
				} else {
					ans = ans + lf*rf
				}
			}
		}
		ans = ans % modulo
		if isTrue {
			dp[i][j][0] = ans
		} else {
			dp[i][j][1] = ans
		}
		return ans
	}
	return solve(s, 0, len(s)-1, true)
}
func getKey(i, j int, isTrue bool) string {
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(i), strconv.Itoa(j), strconv.FormatBool(isTrue))
}
func main() {
	fmt.Println(countWays1(7, "T|T&F^T"))
	fmt.Println(countWays1(5, "T^F|F"))
}
