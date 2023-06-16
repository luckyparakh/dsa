package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	words := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		words[s[i]] += 1
		words[t[i]] -= 1
	}
	// fmt.Println(words)
	for _, v := range words {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isAnagram("", ""))
	fmt.Println(isAnagram("", "a"))
	fmt.Println(isAnagram("abc", "cba"))
}
