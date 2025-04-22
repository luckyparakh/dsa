package main

import (
	"fmt"
	"sort"
)

type item struct {
	cars []int
	time float64
}
type stack struct {
	items []*item
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) top() (*item, bool) {
	if s.isEmpty() {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}
func (s *stack) push(i *item) {
	s.items = append(s.items, i)
}

func (s *stack) pop() (*item, bool) {
	if s.isEmpty() {
		return nil, false
	}
	l := len(s.items)
	i := s.items[l-1]
	s.items = s.items[:l-1]
	return i, true
}
func carFleet(target int, position []int, speed []int) int {
	l := len(position)
	if l == 1 {
		return 1
	}
	m := make(map[int]int)
	for k, v := range position {
		m[v] = speed[k]
	}
	sort.Ints(position)
	s := stack{
		items: []*item{},
	}
	for i := l - 1; i >= 0; i-- {
		t := float64(target - position[i]) / float64(m[position[i]])
		currentItem := item{
			cars: []int{position[i]},
			time: t,
		}
		if v, more := s.top(); more {
			if v.time >= t {
				v.cars = append(v.cars, position[i])
			} else {
				s.push(&currentItem)
			}
		} else {
			s.push(&currentItem)
		}
	}
	for i := range s.items {
		fmt.Println(*(s.items[i]))
	}
	return len(s.items)
}
