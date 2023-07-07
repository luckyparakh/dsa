package main

import "fmt"

// https://leetcode.com/problems/min-stack/
// Next large number, parent problem
// Will use stack as if we use brute force than there will 2 loop and inner loop is dependent on outer loop then it is case 
// of stack.
// Start from Left, store number in stack
// if top of stack is larger it is the answer
// Otherwise keep poping the number till a large number is found or stack gets empty
// Push the number to stack

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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stk := NewStack()
	op := make(map[int]int)
	op1 := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		if stk.isEmpty() {
			op[nums2[i]] = -1
		} else if stk.Top() > nums2[i] {
			op[nums2[i]] = stk.Top()
		} else {
			for stk.Top() < nums2[i] {
				stk.Pop()
				if stk.isEmpty() {
					break
				}
			}
			if stk.isEmpty() {
				op[nums2[i]] = -1
			} else {
				op[nums2[i]] = stk.Top()
			}

		}
		stk.Push(nums2[i])

	}
	for _, num := range nums1 {
		op1 = append(op1, op[num])
	}
	return op1
}
func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
