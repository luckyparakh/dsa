package main

import "fmt"

// https://leetcode.com/problems/word-break-ii/description/
// https://leetcode.com/problems/word-break/description/

func wordBreak(s string, wordDict []string) []string {
	ans := []string{}
	words := make(map[string]bool)
	for _, v := range wordDict {
		if _, ok := words[v]; !ok {
			words[v] = true
		}
	}
	var h func(string, []string)
	h = func(st string, temp []string) {
		if st == "" {
			t := ""
			l := len(temp)
			for k, v := range temp {
				t += v
				if k < l-1 {
					t += " "
				}
			}
			fmt.Println(t)
			ans = append(ans, t)
			return
		}
		for i := 1; i <= len(st); i++ {
			if words[st[:i]] {
				temp = append(temp, st[:i])
				h(st[i:], temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
	h(s, []string{})
	return ans
}
func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
