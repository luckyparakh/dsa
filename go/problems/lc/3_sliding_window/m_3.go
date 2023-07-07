package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

// Maintain a map and put value of rp in map
// If any value of rp is already present in map, then keep rp unchanged and delete all
// elements present between lp and that current rp value. Keep lp incrementing.

func lengthOfLongestSubstring(s string) int {
	lp, rp, l := 0, 0, 0
	mp := make(map[byte]int)
	for rp < len(s) {
		if _, found := mp[s[rp]]; found {
			delete(mp, s[lp])
			lp++
		} else {
			if rp-lp+1 > l {
				l = rp - lp + 1
			}
			mp[s[rp]] = 1
			rp++
		}
	}
	fmt.Println(mp)
	return l
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("abcdefaxyzdf"))
	fmt.Println(lengthOfLongestSubstring("aaaaa"))
	fmt.Println(lengthOfLongestSubstring("a"))
}
