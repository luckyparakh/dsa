package main

// https://www.youtube.com/watch?v=xlKrk3ZO3iM&list=PL_z_8CaSLPWdbOTog8Jxk9XOjzUs3egMP&index=11
// Don't attempt any problem without drawing its recursion diagram
import (
	"fmt"
)

func increasingOrder(numDigit int) []int {
	if numDigit == 0 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	var h func(int, []int)
	arr := []int{}
	ans := []int{}
	h = func(n int, arr []int) {
		if n != 0 {
			arr = append(arr, n)
		}
		if len(arr) == numDigit {
			a := 0
			for i := range arr {
				a = a*10 + arr[i]
			}
			ans = append(ans, a)
			return
		}
		for i := n + 1; i < 10; i++ {
			h(i, arr)
		}
		if n != 0 {
			arr = arr[:len(arr)-1]
		}
	}
	h(0, arr)
	fmt.Println(ans)
	return ans
}

func main() {
	increasingOrder(2)
}
