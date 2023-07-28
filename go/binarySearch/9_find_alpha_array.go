// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	s := 0
	e := len(letters) - 1

	for s <= e {
		m := s + (e-s)/2
		if letters[m] == target {
			s = m + 1
		}
		if letters[m] > target {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	n := (e + 1) % len(letters)
	return letters[n]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
	fmt.Println(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))
	fmt.Println(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'y'))
}
