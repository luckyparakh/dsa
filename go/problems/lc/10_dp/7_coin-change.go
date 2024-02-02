// https://leetcode.com/problems/coin-change/
package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	var driver func([]int, int, int) int
	driver = func(coins []int, amount, element int) int {
		// fmt.Println(amount, element)
		if amount == 0 {
			return 0
		}
		if element == 0 {
			return math.MaxInt
		}
		index := element - 1
		if coins[index] <= amount {
			taken := driver(coins, amount-coins[index], element)
			notTaken := driver(coins, amount, element-1)
			min := taken
			if min > notTaken {
				min = notTaken
			}
			if min == math.MaxInt {
				return math.MaxInt
			}
			return min + 1
		}
		return driver(coins, amount, element-1)
	}
	result := driver(coins, amount, len(coins))
	if result == math.MaxInt {
		return -1
	}
	return result
}

func main() {
	fmt.Println(coinChange([]int{5}, 5))        //1
	fmt.Println(coinChange([]int{3}, 6))        //2
	fmt.Println(coinChange([]int{2, 4}, 11))    //-1
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) //3
}
