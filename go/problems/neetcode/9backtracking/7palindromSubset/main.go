// https://leetcode.com/problems/palindrome-partitioning/description/
package main

import "fmt"

func partition(s string) [][]string {
	ans := [][]string{}
	t := []string{}
	var h func(string, []string)
	h = func(st string, temp []string) {
		if len(st) == 0 {
			a := make([]string, len(temp))
			copy(a, temp)
			ans = append(ans, a)

			// wrong code,temp is directly appended to ans. Since temp is modified in later recursive calls,
			// ans ends up storing references to the same underlying slice instead of different copies.
			// ans = append(ans, temp)
			return
		}

		for j := 1; j <= len(st); j++ {
			st1 := st[:j]
			if isPalindrome(st1) {
				temp = append(temp, st1)
				h(st[j:], temp)
				lt := len(temp)
				temp = temp[:lt-1]
			}
		}
	}
	h(s, t)
	return ans
}
func isPalindrome(st string) bool {
	s, e := 0, len(st)-1
	for s < e {
		if st[s] != st[e] {
			return false
		}
		s++
		e--
	}
	return true
}
func main() {
	fmt.Println(partition("aab"))
}
