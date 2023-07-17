package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	stk := NewStack()
	pairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, v := range s {
		pair, ok := pairs[v]
		if !ok {
			stk.Push(int(v))
		} else {
			if stk.isEmpty() || stk.Top() != int(pair) {
				return false
			} else {
				stk.Pop()
			}
		}
	}
	return stk.isEmpty() == true
}

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items)
	le := s.items[l-1]
	s.items = s.items[:l-1]
	return le
}

func (s *Stack) Top() int {
	l := len(s.items)
	return s.items[l-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid(")"))
	fmt.Println(isValid("("))
	fmt.Println(isValid("(()))"))
}
