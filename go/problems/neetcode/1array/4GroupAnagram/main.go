package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	hm := make(map[string][]string)
	for _, str := range strs {
		sorted := sortStr(str)
		hm[sorted] = append(hm[sorted], str)
	}
	ans := make([][]string, len(hm))
	for _, v := range hm {
		ans = append(ans, v)
	}
	return ans
}

func sortStr(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams1(strs []string) [][]string {
	hm := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for i := range str {
			key[str[i]-'a']++
		}
		hm[key] = append(hm[key], str)
	}
	var ans [][]string
	for _, v := range hm {
		ans = append(ans, v)
	}
	return ans
}
