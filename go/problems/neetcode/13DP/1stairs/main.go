package main

import "fmt"

func climbStairs(n int) int {
	// it will fail for big number like 45 or more
	c := 0
	var helper func(int)
	helper = func(n int) {
		if n < 0 {
			return
		}
		if n == 0 {
			c++
			return
		}
		helper(n - 1)
		helper(n - 2)
	}
	helper(n)
	return c
}
func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStairs2(n int) int {
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

/*
n=1;a=1,b=1
n=2;a=1,b=2
n=3;a=2,b=3
*/