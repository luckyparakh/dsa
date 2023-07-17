package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stk := NewStack()
	for _, v := range tokens {
		// fmt.Println(stk.items)
		if v == "+" || v == "-" || v == "*" || v == "/" {
			if len(stk.items) >= 2 {
				op1 := stk.Pop()
				op2 := stk.Pop()
				switch v {
				case "+":
					stk.Push(op1 + op2)
				case "-":
					stk.Push(op2 - op1)
				case "*":
					stk.Push(op2 * op1)
				case "/":
					stk.Push(op2 / op1)
				}
			}
		} else {
			val, _ := strconv.ParseInt(v, 10, 32)
			stk.Push(int(val))
		}
	}
	if len(stk.items) != 0 {
		return stk.Top()
	}
	return -1
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
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Println(evalRPN([]string{"*"}))
	fmt.Println(evalRPN([]string{"2"}))
}
