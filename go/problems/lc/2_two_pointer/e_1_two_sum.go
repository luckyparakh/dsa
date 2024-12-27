// https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
	"sort"
)

func twoSumAgain(nums []int, target int) bool {
	sort.Ints(nums)
	fmt.Println(nums)
	lp, hp := 0, len(nums)-1
	for lp < hp {
		sum := nums[lp] + nums[hp]
		if sum > target {
			hp--
		} else if sum < target {
			lp++
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(twoSumAgain([]int{11, 2, 15, 7}, 9))
}
