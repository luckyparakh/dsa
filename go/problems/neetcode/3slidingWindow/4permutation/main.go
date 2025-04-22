package main

func checkInclusion(s1 string, s2 string) bool {
	countD, d := make(map[byte]int), make(map[byte]int)
	for r := range s1 {
		countD[s1[r]]++
	}
	start, l, ans := 0, len(s1), false
	for end := range s2 {
		d[s2[end]]++
		for end-start+1 > l {
			d[s2[start]]--
			start++
		}
		if end-start+1 == l {
			a := true
			for i := range countD {
				if countD[i] != d[i] {
					a = false
					break
				}
			}
			if a {
				return a
			}
		}
	}
	return ans
}

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	countD, d := make(map[byte]int), make(map[byte]int)
	for r := range p {
		countD[p[r]]++
	}
	start, l, ans := 0, len(p), []int{}
	for end := range s {
		d[s[end]]++
		for end-start+1 > l {
			d[s[start]]--
			start++
		}
		if end-start+1 == l {
			a := true
			for i := range countD {
				if countD[i] != d[i] {
					a = false
					break
				}
			}
			if a {
				ans = append(ans, start)
			}
		}
	}
	return ans
}
