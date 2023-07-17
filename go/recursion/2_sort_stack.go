package main

import "fmt"

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
func sortStk(s *Stack) *Stack {
	if len(s.items) == 1 {
		return s
	}
	le := s.Pop()
	return insert(sortStk(s), le)
}
func insert(stk *Stack, ele int) *Stack {
	sl := stk.items
	s := 0
	e := len(sl) - 1
	for s <= e {
		m := s + (e-s)/2
		if sl[m] == ele {
			break
		} else if sl[m] > ele {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	position := e + 1
	aI := []int{}
	if position == len(sl) {
		aI = stk.items
		aI = append(aI, ele)
	} else {
		for k, v := range sl {
			if k == position {
				aI = append(aI, ele)
			}
			aI = append(aI, v)
		}
	}
	stk.items = aI
	return stk
}
func main() {
	fmt.Println(sortStk(&Stack{
		items: []int{8, 17, 15, 27, 25, 20},
	}))
	fmt.Println(sortStk(&Stack{
		items: []int{8},
	}))
}
