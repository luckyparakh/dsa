// https://leetcode.com/problems/container-with-most-water/

package main

import "fmt"

// Area is w*h and to get max both w and h should be maximum in value
// w can be max if pointers are far.
// Represent 2 heights as lp and hp, as we need max width initialize lp and hp as 0 and len-1
// h will |height[lp] - height[hp]|
// To keep h maximum move that pointer which points to less value among both the heights
// If both pointers are equal, move any pointer it doesn't matter much

func maxArea(height []int) int {
	lp, hp, max := 0, len(height)-1, 0
	for lp < hp {
		w := hp - lp
		h := height[lp]
		if height[lp] < height[hp] {
			lp++
		} else if height[lp] > height[hp] {
			h = height[hp]
			hp--
		} else {
			lp++
		}
		if w*h > max {
			max = w * h
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 9, 6, 2, 5, 4, 8, 9, 7}))
	fmt.Println(maxArea([]int{1, 1,0,2}))
}
