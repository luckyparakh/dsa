package main

import "fmt"

// https://leetcode.com/problems/permutation-in-string/description/

// Maintain size of window as per match string during that keep on decrease the count of char in map present at right pointer
// To maintain size of window remove char present at left pointer and increase the count of char in map present at left pointer
// If count of any char, reaches zero then it mean in that window that char count matches count with 'match' string and reduce count by one
// If count reached zero then all char of 'match' string is present in the window and mean that permutation is found.

func checkInclusion(match string, permStr string) bool {
	// Edge Case
	// Len of permStr should be equal or less than len of match 
	// as permutation of match can't be greater in size in permStr
	if len(permStr) >= len(match) {
		l, r := 0, 0
		m := make(map[byte]int)

		// Count char in match string and put in map
		for i := range match {
			m[match[i]]++
		}
		// Initialize c with len of map which denotes number of distinct characters in match string
		c := len(m)

		for r < len(permStr) {
			if _, found := m[permStr[r]]; found {
				m[permStr[r]]--
				if m[permStr[r]] == 0 {
					c--
				}
			}
			if c == 0 {
				return true
			}
			if r-l+1 < len(match) {
				r++
				continue
			}
			if _, found := m[permStr[l]]; found {
				if m[permStr[l]] == 0 {
					c++
				}
				m[permStr[l]]++
			}
			l++
			r++
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab","eidbaooo"))
	fmt.Println(checkInclusion("ab","eidbcooo"))
	fmt.Println(checkInclusion("ab","ba"))
	fmt.Println(checkInclusion("abc","ba"))
	fmt.Println(checkInclusion("aa","aaaabaaaaaca"))
}
