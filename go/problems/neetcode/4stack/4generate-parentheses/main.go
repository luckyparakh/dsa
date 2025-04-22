package main

import "fmt"

func generateParenthesis(n int) []string {
	l := n * 2
	r := []rune{'('}
	d := make(map[rune]int)
	d['('] = 1
	ans := []string{}
	var h func([]rune, int)
	h = func(r []rune, start int) {
		if start == l {
			if d['('] == n && d[')'] == n {
				ans = append(ans, string(r))
				return
			}
		}
		if d['('] <= n && d[')'] <= d['('] {
			r = append(r, '(')
			d['(']++
			h(r, start+1)
			// backtrack
			r = r[:len(r)-1]
			d['(']--
			r = append(r, ')')
			d[')']++
			h(r, start+1)
			r = r[:len(r)-1]
			d[')']--
		} else {
			return
		}
	}
	h(r, 1)
	return ans
}
func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(2))
}
