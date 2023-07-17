package main

// https://www.youtube.com/watch?v=8YXQ68oHjAs&list=PL_z_8CaSLPWeT1ffjiImo0sYTcnLzo-wY&index=9
// use Base, Hypothesis and Induction method
// Reduce input

import "fmt"

type Stack struct {
	items []int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(ele int) {
	s.items = append(s.items, ele)
}
func (s *Stack) Pop() int {
	leIndex := len(s.items) - 1
	le := s.items[leIndex]
	s.items = s.items[:leIndex]
	return le
}
func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack) Top() int {
	return s.items[0]
}
func (s *Stack) Len() int {
	return len(s.items)
}
func insertAtEnd(stk *Stack, ele int) *Stack {
	if stk.Len() == 0 {
		stk.Push(ele)
		return stk
	}
	le := stk.Pop()
	stk = insertAtEnd(stk, ele)
	stk.Push(le)
	return stk
}
func reverseStack(stk *Stack) *Stack {
	if stk.Len() == 1 || stk.Len() == 0 {
		return stk
	}
	le := stk.Pop()
	stk = reverseStack(stk)
	return insertAtEnd(stk, le)
}
func main() {
	fmt.Println(reverseStack(&Stack{
		items: []int{8},
	}))
	fmt.Println(reverseStack(&Stack{
		items: []int{10, 8},
	}))
	fmt.Println(reverseStack(&Stack{
		items: []int{},
	}))
	fmt.Println(reverseStack(&Stack{
		items: []int{8, 17, 15, 27, 25, 20},
	}))
}
