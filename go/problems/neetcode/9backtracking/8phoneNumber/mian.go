package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	ans := []string{}
	digitLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var h func(int, []byte)
	h = func(indexOfDigit int, t []byte) {
		if indexOfDigit == len(digits) {
			ans = append(ans, string(t))
			return
		}
		letters := digitLetters[digits[indexOfDigit]]
		for j := 0; j < len(letters); j++ {
			t = append(t, letters[j])
			// fmt.Println(t)
			h(indexOfDigit+1, t)
			t = t[:len(t)-1]
			// fmt.Println(t)
		}
	}
	h(0, []byte{})
	return ans
}
func main() {
	fmt.Println(letterCombinations("234"))
}
