package main
// https://leetcode.com/problems/valid-palindrome/description/
import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(",."))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	fp, ep := 0, len(s)-1
	exp:=regexp.MustCompile("[a-z0-9]")
	for fp < ep {
		if !exp.MatchString(string(s[fp])) {
			fp++
			continue
		}
		if !exp.MatchString(string(s[ep])) {
			ep--
			continue
		}
		if s[fp] != s[ep] {
			return false
		}
		fp++
		ep--
	}
	return true
}
