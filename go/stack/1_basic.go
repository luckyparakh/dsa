package main

import (
	"sync"
)

type Item interface{}
type Stack struct {
	items []Item
	mutex sync.Mutex
}

func newEmptyStack() *Stack {
	return &Stack{}
}

func newStack(items []Item) *Stack {
	return &Stack{
		items: items,
	}
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item Item) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, item)
}

func (s *Stack) Pop(item Item) Item {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.isEmpty() {
		return nil
	}
	lastItem := s.items[len(s.items)-1]
	s.items = s.items[:s.Len()-1]
	return lastItem
}

func (s *Stack) Top() Item {
	if s.isEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

func main() {

}
