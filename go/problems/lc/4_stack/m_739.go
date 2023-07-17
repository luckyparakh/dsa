package main

import "fmt"

type item struct {
	val      int
	position int
}

func dailyTemperatures(temperatures []int) []int {
	ans := []int{}
	ngl := nearestGreatLeft(temperatures)
	for k := range temperatures {
		ans = append(ans, ngl[k]-k)
	}
	return ans
}
func nearestGreatLeft(temperatures []int) []int {
	s := make([]int, len(temperatures))
	stk := []item{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stk) != 0 {
			if temperatures[i] >= stk[len(stk)-1].val {
				stk = stk[:len(stk)-1]
			} else {
				s[i] = stk[len(stk)-1].position
				stk = append(stk, item{
					val:      temperatures[i],
					position: i,
				})
				break
			}
		}
		if len(stk) == 0 {
			stk = append(stk, item{
				val:      temperatures[i],
				position: i,
			})
			s[i] = i
		}
	}
	fmt.Println(s)
	return s
}

// https://leetcode.com/problems/daily-temperatures/
func main() {
	// fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}
