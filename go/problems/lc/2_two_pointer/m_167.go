// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
package main

import "fmt"

// Start LP from 0 and HP from length -1 of slice.
// Now take of sum of a[lp]+a[hp].
// Since slice is sorted sum is greater than reduce HP or is lesser than increase LP.
// If sum is equal then print lp and hp

func twoSum(numbers []int, target int) []int {
	lp, hp := 0, len(numbers)-1

	for lp < hp {
		if numbers[lp]+numbers[hp] < target {
			lp++
			continue
		} else if numbers[lp]+numbers[hp] > target {
			hp--
			continue
		} else {
			return []int{lp + 1, hp + 1}
		}
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{-1, 0}, -1))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{1, 1, 1, 2, 3, 4}, 2))
	fmt.Println(twoSum([]int{1, 1, 1, 2, 3, 4}, 3))
}
