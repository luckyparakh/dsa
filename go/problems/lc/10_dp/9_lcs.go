// https://leetcode.com/problems/longest-common-subsequence/

package main

import (
	"fmt"
	"math"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	var driver func(string, string, int, int) int
	driver = func(s1, s2 string, i1, i2 int) int {
		if i1 == 0 || i2 == 0 {
			return 0
		}
		if s1[i1-1] == s2[i2-1] {
			return 1 + driver(s1, s2, i1-1, i2-1)

		} else {
			return int(math.Max(float64(driver(s1, s2, i1, i2-1)),
				float64(driver(s1, s2, i1-1, i2))))
		}
	}
	return driver(text1, text2, len(text1), len(text2))
}
func longestCommonSubsequenceMemo(text1 string, text2 string) int {
	i1, i2 := len(text1), len(text2)
	t := make([][]int, i1+1)
	for i := 0; i < i1+1; i++ {
		s := make([]int, i2+1)
		for j := 0; j < i2+1; j++ {
			s[j] = -1
		}
		t[i] = s
	}

	var driver func(string, string, int, int) int
	driver = func(s1, s2 string, i1, i2 int) int {
		if i1 == 0 || i2 == 0 {
			return 0
		}
		if t[i1][i2] != -1 {
			return t[i1][i2]
		}
		if s1[i1-1] == s2[i2-1] {
			t[i1][i2] = 1 + driver(s1, s2, i1-1, i2-1)
			return t[i1][i2]
		} else {
			t[i1][i2] = int(math.Max(float64(driver(s1, s2, i1, i2-1)),
				float64(driver(s1, s2, i1-1, i2))))
			return t[i1][i2]
		}
	}
	return driver(text1, text2, len(text1), len(text2))
}

func longestCommonSubsequenceIter(text1 string, text2 string) int {
	i1, i2 := len(text1), len(text2)
	t := make([][]int, i1+1)
	for i := 0; i < i1+1; i++ {
		s := make([]int, i2+1)
		for j := 0; j < i2+1; j++ {
			s[j] = -1
		}
		t[i] = s
	}
	for i := 0; i < i1+1; i++ {
		for j := 0; j < i2+1; j++ {
			if i == 0 || j == 0 {
				t[i][j] = 0
			} else {
				if text1[i-1] == text2[j-1] {
					t[i][j] = 1 + t[i-1][j-1]
				} else {
					t[i][j] = int(math.Max(float64(t[i][j-1]),
						float64(t[i-1][j])))
				}
			}
		}
	}
	
	return t[i1][i2]
}
func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ade")) //3
	fmt.Println(longestCommonSubsequence("abc", "abc"))   //3
	fmt.Println(longestCommonSubsequence("abc", "def"))   //0
	fmt.Println(longestCommonSubsequenceIter("abcde", "ade")) //3
	fmt.Println(longestCommonSubsequenceIter("abc", "abc"))   //3
	fmt.Println(longestCommonSubsequenceIter("abc", "def"))   //0
	fmt.Println(longestCommonSubsequenceIter("abac", "cab"))   //0
}
