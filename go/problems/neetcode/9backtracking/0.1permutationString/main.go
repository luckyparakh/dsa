package main

import (
	"fmt"
	"strings"
)

func permutationString(s string) {
	var helper func(s *string, i int)
	ans := []string{}
	d := make(map[string]bool)
	helper = func(s *string, start int) {
		if start == len(*s)-1 {
			if !d[*s] {
				d[*s] = true
				ans = append(ans, *s)
			}
			return
		}
		for j := start; j < len(*s); j++ {
			// fmt.Println(*s)
			swap(s, start, j)
			// fmt.Println("after swap", *s, j)
			helper(s, start+1)
			swap(s, start, j)
		}
		// }
	}
	helper(&s, 0)
	fmt.Println(ans)
}
func swap(s *string, i, j int) {
	sArr := strings.Split(*s, "")
	sArr[i], sArr[j] = sArr[j], sArr[i]
	*s = strings.Join(sArr, "")
}
func permutationStringBetter(s string) {
	r := []rune(s)
	ans := []string{}
	var h func([]rune, int)
	h = func(r []rune, start int) {
		l := len(r)
		if start == l-1 {
			ans = append(ans, string(r))
			return
		}
		for j := start; j < l; j++ {
			if start != j && r[start] == r[j] {
				continue
			}
			r[start], r[j] = r[j], r[start]
			h(r, start+1)
			// Back Track
			r[start], r[j] = r[j], r[start]
		}
	}
	h(r, 0)
	fmt.Println(ans)
}

func main() {
	permutationString("aabc")
	permutationString("abc")

	permutationStringBetter("aabc")
	permutationStringBetter("abc")
}
