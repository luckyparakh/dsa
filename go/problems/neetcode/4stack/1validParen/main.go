package main

import "fmt"

type Stack struct {
	items []rune
}

func (s *Stack) push(val rune) {
	s.items = append(s.items, val)
}
func (s *Stack) pop() (rune, error) {
	l := len(s.items)
	if l <= 0 {
		return 0, fmt.Errorf("empty stack")
	}
	v := s.items[l-1]
	s.items = s.items[:l-1]
	return v, nil
}
func (s *Stack) top() (rune, error) {
	l := len(s.items)
	if l <= 0 {
		return 0, fmt.Errorf("empty stack")
	}
	return s.items[l-1], nil
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	elements := make(map[rune]rune)
	elements[')'] = '('
	elements[']'] = '['
	elements['}'] = '{'
	st := Stack{items: []rune{}}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st.push(v)
			// fmt.Println(st.items)
		} else {
			// fmt.Println(st.items)
			if t, err := st.pop(); err != nil {
				return false
			} else {
				if t != elements[v] {
					return false
				}
			}
		}
	}
	// fmt.Println(st.items)
	return len(st.items) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid(""))
	fmt.Println(isValid("}"))
	fmt.Println(isValid("({()}"))
}
