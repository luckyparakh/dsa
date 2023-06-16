package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams1(strs []string) [][]string {
	// Logic: Loop through slice of words; o[n2]
	// Now compare each word against key of map, if key and word is anagram save word in values of that key
	// If word is not a anagram to any key in map, save word as key of map
	// Finally print values
	anagrams := make(map[string][]string)
	for _, str := range strs {
		anagramFound := false
		for k, v := range anagrams {
			if isAnagram(k, str) {
				anagramFound = true
				anagrams[k] = append(v, str)
				break
			}
		}
		if !anagramFound {
			anagrams[str] = append(anagrams[str], str)
		}
	}
	fmt.Println(anagrams)
	output := [][]string{}
	for _, v := range anagrams {
		output = append(output, v)
	}
	return output
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	words := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		words[s[i]] += 1
		words[t[i]] -= 1
	}
	for _, v := range words {
		if v != 0 {
			return false
		}
	}
	return true
}
func groupAnagrams(strs []string) [][]string {
	// Better Way (o[n])
	// Logic: Loop through slice of words
	// Sort each word and compare them with key present in map, if present then append in key's value
	// And if word is not present in map then add in map
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sorted(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	// fmt.Println(anagrams)
	output := [][]string{}
	for _, v := range anagrams {
		output = append(output, v)
	}
	return output
}
func sorted(str string) string {
	y := strings.Split(str, "")
	sort.Strings(y)
	return strings.Join(y, "")
}

func main() {
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{"eat", "ate", "tan", "man", "bat", "tab"}))

}
