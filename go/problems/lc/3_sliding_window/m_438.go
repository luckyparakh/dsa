package main

import "fmt"

// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

func findAnagrams(s string, p string) []int {
	ans := []int{}
	if len(p) >= len(s) {
		l, r := 0, 0
		m := make(map[byte]int)

		// Count char in match string and put in map
		for i := range s {
			m[s[i]]++
		}
		// Initialize c with len of map which denotes number of distinct characters in match string
		c := len(m)

		for r < len(p) {
			if _, found := m[p[r]]; found {
				m[p[r]]--
				if m[p[r]] == 0 {
					c--
				}
			}
			if c == 0 {
				ans = append(ans, l)
			}
			if r-l+1 < len(s) {
				r++
				continue
			}
			if _, found := m[p[l]]; found {
				if m[p[l]] == 0 {
					c++
				}
				m[p[l]]++
			}
			l++
			r++
		}
		return ans
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("ab", "eidbaooo"))
	fmt.Println(findAnagrams("ab", "eidbcooo"))
	fmt.Println(findAnagrams("ab", "ba"))
	fmt.Println(findAnagrams("abc", "ba"))
	fmt.Println(findAnagrams("ab", "abab"))
	fmt.Println(findAnagrams("aa", "aaaabaaaaaca"))
}
