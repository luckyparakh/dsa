package main

import "fmt"

type item struct {
	val   int
	index int
}
type stack struct {
	items []item
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := stack{
		items: []item{
			item{val: temperatures[0], index: 0},
		},
	}

	for i := 1; i < len(temperatures); i++ {
		for val, err := st.top(); val < temperatures[i] && err == nil; val, err = st.top() {
			item, _ := st.pop()
			idx := (*item).index
			ans[idx] = i - idx
		}
		st.push(item{val: temperatures[i], index: i})
	}
	return ans
}
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func (s *stack) push(val item) {
	s.items = append(s.items, val)
}
func (s *stack) top() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("")
	}
	return s.items[len(s.items)-1].val, nil
}

func (s *stack) pop() (*item, error) {
	l := len(s.items)
	if l == 0 {
		return nil, fmt.Errorf("")
	}
	t := s.items[l-1]
	s.items = s.items[:l-1]
	return &t, nil
}
