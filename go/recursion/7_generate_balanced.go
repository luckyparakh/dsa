// https://leetcode.com/problems/generate-parentheses/

package main

import "fmt"

func generateParenthesis(n int) []string {
	d := make(map[byte]int)
	d['('] = n - 1
	d[')'] = n
	op := []string{}
	gp("(", d, &op)
	return op
}
func gp(s string, d map[byte]int, op *[]string) {
	// fmt.Println("new call")
	if d['('] == 0 {
		for i := 1; i <= d[')']; i++ {
			s = fmt.Sprintf("%s%s", s, ")")
		}
		*op = append(*op, s)
		// fmt.Println("final",op, d)
		return
	}
	last := s[len(s)-1]
	// fmt.Println("last:",last)
	// fmt.Println("input dict:",d)
	if last == '(' {
		if d['('] > 0 {

			// fmt.Println("1", s)
			d1 := make(map[byte]int)
			for k, v := range d {
				d1[k] = v
			}
			d1['('] = d1['('] - 1
			gp(fmt.Sprintf("%s%s", s, "("), d1, op)
		}
		if d[')'] > 0 {
			// fmt.Println("2", s)
			d2 := make(map[byte]int)
			for k, v := range d {
				d2[k] = v
			}

			d2[')'] = d2[')'] - 1
			gp(fmt.Sprintf("%s%s", s, ")"), d2, op)
		}
	}
	if last == ')' {
		if d['('] > 0 {
			// fmt.Println("3", s)
			d3 := make(map[byte]int)
			for k, v := range d {
				d3[k] = v
			}
			d3['('] = d3['('] - 1
			gp(fmt.Sprintf("%s%s", s, "("), d3, op)
		}
		if d[')'] > d['('] {

			// fmt.Println("4", s)
			d4 := make(map[byte]int)
			for k, v := range d {
				d4[k] = v
			}
			d4[')'] = d4[')'] - 1
			gp(fmt.Sprintf("%s%s", s, ")"), d4, op)
		}
	}
}
func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
