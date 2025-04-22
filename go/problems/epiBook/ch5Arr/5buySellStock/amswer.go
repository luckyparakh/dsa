// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
// Intuition is to find lowest valley and then take diff at current to find max profit
package main

func maxProfit(prices []int) int {
	minPriceTillThisINdex := 100000
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPriceTillThisINdex {
			minPriceTillThisINdex = prices[i]
		}
		if prices[i]-minPriceTillThisINdex > maxProfit {
			maxProfit = prices[i] - minPriceTillThisINdex
		}
	}
	return maxProfit
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// https://algo.monster/liteproblems/122
func maxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
