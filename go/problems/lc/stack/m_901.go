package main

import "fmt"

// https://leetcode.com/problems/online-stock-span/

// Similar to find nearest large number
// Find nearest large number along with that in stack save that large number with its span day, push that number in stack
// Check top, if number is large then return span day as 1
// If number is small or equal, then pop that number from stack and add span day of that number to current span day sum
// and stop at large number and push that number in stack


type Stack struct {
	// []int slice will store price and span time
	items [][]int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(i []int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() []int {
	l := len(s.items)
	le := s.items[l-1]
	s.items = s.items[:l-1]
	return le
}

func (s *Stack) Top() []int {
	l := len(s.items)
	return s.items[l-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

type StockSpanner struct {
	s Stack
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	sum := 1
	for !this.s.isEmpty() {
		if this.s.Top()[0] > price {
			this.s.Push([]int{price, sum})
			return sum
		}
		sum += this.s.Pop()[1]
	}
	this.s.Push([]int{price, sum})
	return sum

}

func main() {
	stockSpanner := Constructor()
	fmt.Println(stockSpanner.Next(100))
	fmt.Println(stockSpanner.Next(80))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(70))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(75))
	fmt.Println(stockSpanner.Next(85))
	fmt.Println(stockSpanner.Next(100))
	fmt.Println(stockSpanner.Next(100))
	fmt.Println(stockSpanner.Next(100))
}
