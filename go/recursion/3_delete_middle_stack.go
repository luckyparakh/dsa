package main

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
func deleteMiddle(stk *Stack) *Stack {
	l := len(stk.items)
	if l == 0 {
		return stk
	}
	middle := l/2 + 1
	popCount := l - middle
	return deleteM(stk, popCount)
}
func deleteM(stk *Stack, pc int) *Stack {
	if pc == 0 {
		stk.Pop()
		return stk
	}
	ele := stk.Pop()
	deleteM(stk, pc-1)
	stk.Push(ele)
	return stk
}
func main() {
	fmt.Println(deleteMiddle(&Stack{
		items: []int{8},
	}))
	fmt.Println(deleteMiddle(&Stack{
		items: []int{10, 8},
	}))
	fmt.Println(deleteMiddle(&Stack{
		items: []int{},
	}))
	fmt.Println(deleteMiddle(&Stack{
		items: []int{8, 17, 15, 27, 25, 20},
	}))
	fmt.Println(deleteMiddle(&Stack{
		items: []int{8, 17, 27, 25, 20},
	}))
}
