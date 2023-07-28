package main

import "fmt"

// https://leetcode.com/problems/find-the-winner-of-the-circular-game/description/
// https://www.youtube.com/watch?v=ULUNeD0N9yI&list=PL_z_8CaSLPWeT1ffjiImo0sYTcnLzo-wY&index=19
func findTheWinner(n int, k int) int {
	index := []int{}
	for i := 1; i <= n; i++ {
		index = append(index, i)
	}
	return solve(index, k, 0)
}
func solve(a []int, k, start int) int {
	if len(a) == 1 {
		return a[0]
	}
	newA := []int{}
	toRemove := start + k - 1
	toRemove = toRemove % len(a)
	newA = append(newA, a[:toRemove]...)
	newA = append(newA, a[toRemove+1:]...)
	return solve(newA, k, toRemove)
}
func findTheWinnerLinear(n int, k int) int {
	index := []int{}
	for i := 1; i <= n; i++ {
		index = append(index, i)
	}
	start := 0
	for len(index) != 1 {
		toRemove := (start + k - 1) % len(index)
		index = append(index[:toRemove], index[toRemove+1:]...)
		start = toRemove
	}
	return index[0]
}
func main() {
	fmt.Println(findTheWinner(1, 1))
	fmt.Println(findTheWinner(5, 2))
	fmt.Println(findTheWinner(6, 5))
	fmt.Println(findTheWinner(40, 7))
	fmt.Println(findTheWinnerLinear(1, 1))
	fmt.Println(findTheWinnerLinear(5, 2))
	fmt.Println(findTheWinnerLinear(6, 5))
	fmt.Println(findTheWinnerLinear(40, 7))
}
