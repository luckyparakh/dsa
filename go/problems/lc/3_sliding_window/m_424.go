package main

import "fmt"

// https://leetcode.com/problems/longest-repeating-character-replacement/description/

func characterReplacement(s string, k int) int {
	lp, rp, maxLen := 0, 0, 0
	m := make(map[byte]int)
	for rp < len(s) {
		m[s[rp]]++
		for (rp-lp+1)-findMaxFrequentCharacter(m) > k {
			m[s[lp]]--
			lp++
		}
		if rp-lp+1 > maxLen {
			maxLen = rp - lp + 1
		}
		rp++
	}
	return maxLen
}
func findMaxFrequentCharacter(m map[byte]int) int {
	maxFreq := 0
	for _, v := range m {
		if v > maxFreq {
			maxFreq = v
		}
	}
	return maxFreq
}
func main() {
	fmt.Println(characterReplacement("aabaccacc", 2))
	fmt.Println(characterReplacement("abaa", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
	fmt.Println(characterReplacement("aaa", 2))
	fmt.Println(characterReplacement("aaabb", 0))
	fmt.Println(characterReplacement("a", 0))
}
