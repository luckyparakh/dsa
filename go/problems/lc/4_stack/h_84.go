package main

import "fmt"

// https://leetcode.com/problems/largest-rectangle-in-histogram/
// Before this understand NSR and NSL of stack

func largestRectangleArea(heights []int) int {
	l := nsl(heights)
	r := nsr(heights)
	max := 0
	for k, v := range heights {
		area := v * (r[k] - l[k] - 1)
		if area > max {
			max = area
		}
	}
	return max
}
func nsl(h []int) []int {
	stk := newStack()
	op := make([]int, len(h))
	for k, v := range h {
		if stk.isEmpty() {
			op[k] = -1
			stk.push(item{
				value:    v,
				position: k,
			})
		} else {
			for !stk.isEmpty() {
				if stk.top().value >= v {
					stk.pop()
				} else {
					break
				}
			}
			if stk.isEmpty() {
				op[k] = -1
			} else {
				op[k] = stk.top().position
			}
			stk.push(item{
				value:    v,
				position: k,
			})
		}
	}
	return op
}
func nsr(h []int) []int {
	stk := newStack()
	op := make([]int, len(h))
	for i := len(h) - 1; i >= 0; i-- {
		if stk.isEmpty() {
			op[i] = len(h)
			stk.push(item{
				value:    h[i],
				position: i,
			})
		} else {
			for !stk.isEmpty() {
				if stk.top().value >= h[i] {
					stk.pop()
				} else {
					break
				}
			}
			if stk.isEmpty() {
				op[i] = len(h)
			} else {
				op[i] = stk.top().position
			}

			stk.push(item{
				value:    h[i],
				position: i,
			})
		}
	}
	return op
}

type item struct {
	value    int
	position int
}
type Stack struct {
	h []item
}

func newStack() *Stack {
	return &Stack{}
}
func (s *Stack) isEmpty() bool {
	return len(s.h) == 0
}
func (s *Stack) push(i item) {
	s.h = append(s.h, i)
}
func (s *Stack) pop() item {
	item := s.h[len(s.h)-1]
	s.h = s.h[:len(s.h)-1]
	return item
}
func (s *Stack) top() item {
	return s.h[len(s.h)-1]
}
func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{0, 4}))
}
