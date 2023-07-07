// https://leetcode.com/problems/valid-palindrome/description/
package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	lp := 0
	hp := len(s) - 1
	exp := regexp.MustCompile("[a-z0-9]")
	for lp < hp {
		if !exp.MatchString(string(s[lp])) {
			lp++
			continue
		}
		if !exp.MatchString(string(s[hp])) {
			hp--
			continue
		}
		if s[lp] != s[hp] {
			return false
		}
		lp++
		hp--
	}
	return true
}

func isPalindrome1(s string) bool {
	s = strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) || !unicode.IsLetter(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	lp := 0
	hp := len(s) - 1
	for lp < hp {
		if s[lp] != s[hp] {
			return false
		}
		lp++
		hp--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("a car"))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome(" a    "))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome(".,"))
	fmt.Println(isPalindrome("malayalam"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("0p"))
	fmt.Println(rune(97))
}
