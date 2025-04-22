// https://leetcode.com/problems/jump-game/description/
// https://leetcode.com/problems/jump-game-ii/
package main

import (
	"fmt"
	"math"
)

/*
Intuition in Simple Words:

Imagine you're exploring a maze. You're not just taking the biggest turn at every junction. Instead, you're looking at all the paths you've explored so far and keeping track of the furthest point you've ever reached. If, at any point, you find that you can't go any further from the furthest point you've reached, you know you're trapped. But if you manage to explore a path that takes you to or beyond the exit, you've solved the maze!

Why This Works:

By keeping track of the furthest_reach_so_far, you're essentially simulating all possible jump combinations implicitly. You don't need to explicitly try every path.
If you can reach the end, it means there must be some path (a series of jumps) that allows you to reach or go beyond the last index.
If you can't reach the end, it means no matter which jumps you take, you'll always get stuck before the end.
*/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxSteps := 0
	for i := range nums {
		// Check reachability, maxSteps represent how many max steps are needed to reach here,
		// but if its less than current index which mean from previous indexes we can't reach here
		if maxSteps < i {
			return false
		}
		maxSteps = int(math.Max(float64(maxSteps), float64(i+nums[i])))
	}
	return true
}

func minSteps(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	c, maxSteps, lastStep := 0, 0, 0
	// Iterate over the list excluding the last element as it's unnecessary to jump from the last position.
	for i := 0; i < len(nums)-1; i++ {
		if maxSteps < i {
			return -1
		}
		maxSteps = max(maxSteps, i+nums[i])
		// If we have reached the furthest point to which we had jumped previously,
		// Increment the jump count and update the last reached position to the current max_reach.
		if lastStep == i {
			fmt.Println(i)
			c += 1
			lastStep = maxSteps
		}

	}
	return c
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	fmt.Println(minSteps([]int{2, 3, 1, 1, 4}))
}
