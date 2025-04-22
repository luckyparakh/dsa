// https://leetcode.com/problems/plus-one/description/
// https://gemini.google.com/share/5587532a8699
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(multiply([]int{1, 0}, []int{2, 0}))
	fmt.Println(multiply([]int{1, 2, 3}, []int{-9, 8, 7}))
}

// A brute-force approach might be to convert the array of digits to the equivalent integer,
// increment that, and then convert the resulting value back to an array of digits. For example, if the
// array is <1,2,9>, we would derive the integer 129, add one to get 130, then extract its digits to form
// (1,3,0). When implemented in a language that imposes a limit on the range of values an integer
// type can take, this approach will fail on inputs that encode integers outside of that range.
// We can avoid overflow issues by operating directly on the array of digits. Specifically, we mimic
// the grade-school algorithm for adding integers, which entails adding digits starting from the least
// significant digit, and propagate carries. If the result has an additional digit, e.g., 99 + 1. = 100, there
// is not enough storage in the array for the result-we need three digits to represent 100, but the input
// has only two digits.

func plusOne(digits []int) []int {
	last := len(digits) - 1
	digits[last] += 1

	for i := last; i > 0; i-- {
		if digits[i] != 10 {
			break
		}
		digits[i] = 0
		digits[i-1] += 1
	}
	if digits[0] == 10 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}

func multiply(num1, num2 []int) []int {
	// find sign
	sign := 1
	if num1[0]*num2[0] < 0 {
		sign = -1
	}
	num1[0], num2[0] = int(math.Abs(float64(num1[0]))), int(math.Abs(float64(num2[0])))
	l1, l2 := len(num1), len(num2)
	answer := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			// let say we want ot multiply 54 * 12, so least significant digit calculate by 4*2 will be at end of array.
			// end of array will be sum of len of num1 & num2
			answer[j+i+1] += num1[i] * num2[j]
			// then extra carry is forwarded to digit just before , which will be in i+j in our case
			answer[j+i] += answer[j+i+1] / 10
			answer[j+i+1] = answer[j+i+1] % 10
		}
	}
	fmt.Println(answer)
	indexSlice := 0
	for k, v := range answer {
		if v != 0 {
			indexSlice = k
			break
		}
	}
	answer = answer[indexSlice:]
	answer[0] = sign * answer[0]
	return answer
}
