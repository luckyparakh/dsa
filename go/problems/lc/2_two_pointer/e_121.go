package main

import "fmt"

// Take 2 pointer buy and sell. To maximize profit value of buy should be less than sell in such a way that diff is max.
// Move buy pointer if value of sell pointer is less than buy value
// Also calculate diff meanwhile and store max value in variable

func maxProfit(prices []int) int {
	b, s, max := 0, 0, 0
	for s < len(prices) {
		if prices[b] > prices[s] {
			b = s
		}
		diff := prices[s] - prices[b]
		if diff > max {
			max = diff
		}
		s++
	}
	return max
}
func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7}))
}
