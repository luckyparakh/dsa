package main

import (
	"fmt"
	"strconv"
)
// for DP or BT draw recursive diagram always.
// To find that it is greedy or BT, try to find a counter example for greedy
func LargestNumberKSwaps(n, k int) int {
	r := []rune(fmt.Sprintf("%d", n))
	ans, start := 0, 0
	var h func([]rune, int, int)
	h = func(r []rune, start, k int) {
		if start == len(r)-1 || k == 0 {
			t, _ := strconv.Atoi(string(r))
			if t > ans {
				ans = t
			}
			return
		}
		for i := start; i < len(r); i++ {
			if r[start] > r[i] {
				continue
			}
			r[start], r[i] = r[i], r[start]
			h(r, start+1, k-1)
			r[start], r[i] = r[i], r[start]
		}
	}
	h(r, start, k)
	return ans
}

func main() {
	fmt.Println(LargestNumberKSwaps(4577,3))
	fmt.Println(LargestNumberKSwaps(1234567,4))
	fmt.Println(LargestNumberKSwaps(7654321,4))
	fmt.Println(LargestNumberKSwaps(3435335,3))
	fmt.Println(LargestNumberKSwaps(1034,2))
}
