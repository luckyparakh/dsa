package main

import (
	"fmt"
	"strconv"
)

type Stack []string

func (s *Stack) push(v string) {
	*s = append(*s, v)
}

func (s *Stack) pop() (int, error) {
	var v int
	l := len(*s)
	if l > 0 {
		v, err := strconv.Atoi((*s)[l-1])
		if err != nil {
			return v, err
		}
		(*s) = (*s)[:l-1]
		return v, nil
	}
	return v, fmt.Errorf("invalid operation")
}

func evalRPN(tokens []string) int {
	st := Stack{}
	for _, v := range tokens {
		fmt.Println(v)
		switch v {
		case "+":
			o1, _ := st.pop()
			o2, _ := st.pop()
			st.push(strconv.Itoa(o1 + o2))
		case "-":
			o1, _ := st.pop()
			o2, _ := st.pop()
			st.push(strconv.Itoa(o1 - o2))
		case "*":
			o1, _ := st.pop()
			o2, _ := st.pop()
			st.push(strconv.Itoa(o1 * o2))
		case "/":
			o1, _ := st.pop()
			o2, _ := st.pop()
			st.push(strconv.Itoa(o2 / o1))
		default:
			st.push(v)
		}
		fmt.Println(st)
	}
	fmt.Println(st)
	ans, _ := strconv.Atoi(st[0])
	return ans
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
