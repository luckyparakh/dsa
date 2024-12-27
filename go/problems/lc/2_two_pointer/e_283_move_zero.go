package main

// https://leetcode.com/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	lp, hp := 0, 0
	for hp < len(nums) {
		if nums[lp] != 0 {
			lp++
		} else if nums[lp] == 0 && nums[hp] != 0 {
			nums[lp], nums[hp] = nums[hp], nums[lp]
			lp++
		}
		hp++
	}
}
