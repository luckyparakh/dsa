// https://leetcode.com/problems/shortest-common-supersequence/
// https://www.youtube.com/watch?v=823Grn4_dCQ&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=24
// https://www.youtube.com/watch?v=pHXntFeu6f8
package main

import (
	"fmt"
	"math"
	"strings"
)

func shortestCommonSupersequence(str1 string, str2 string) string {
	lcs := getLcs(str1, str2)
	fmt.Println(lcs)
	ans := []string{}
	s1 := 0
	s2 := 0
	for i := 0; i < len(lcs); i++ {
		for str1[s1] != lcs[i] {
			ans = append(ans, string(str1[s1]))
			s1 = s1 + 1
		}
		for str2[s2] != lcs[i] {
			ans = append(ans, string(str2[s2]))
			s2 = s2 + 1
		}
		ans = append(ans, string(lcs[i]))
		s1 = s1 + 1
		s2 = s2 + 1
		// fmt.Println(ans)
	}
	// add left over string from s1 and s2
	if s1 <= len(str1)-1 {
		ans = append(ans, str1[s1:])
	}
	if s2 <= len(str2)-1 {
		ans = append(ans, str2[s2:])
	}
	return strings.Join(ans, "")
}

func getLcs(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	t := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		s := make([]int, l2+1)
		t[i] = s
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if str1[i-1] == str2[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = int(math.Max(float64(t[i-1][j]), float64(t[i][j-1])))
			}
		}
	}
	fmt.Println(t)

	//find LCS string
	// sb := strings.Builder{}
	lcsLen := t[l1][l2]
	s := make([]string, lcsLen)
	if t[l1][l2] != 0 {
		c := 1
		i := l1
		j := l2
		for i > 0 && j > 0 {
			if str1[i-1] == str2[j-1] {
				s[lcsLen-c] = string(str1[i-1])
				// sb.WriteRune(rune(str1[i-1]))
				i, j = i-1, j-1
				c++
			} else {
				if t[i][j-1] > t[i-1][j] {
					j = j - 1
				} else {
					i = i - 1
				}
			}
		}
		// fmt.Println(s)
	}

	return strings.Join(s, "")
}

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cab"))
	fmt.Println(shortestCommonSupersequence("bbbaaaba", "bbababbb"))
}
