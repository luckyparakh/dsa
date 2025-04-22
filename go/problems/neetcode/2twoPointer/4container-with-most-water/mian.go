package main

func maxArea(nums []int) int {
	l := len(nums)
	s, e, ans := 0, l-1, 0
	// To maximize the area, at each step, we move the pointer corresponding to the shorter line. 
	// This is because moving the shorter line gives us a chance of increasing the height, whereas moving the taller line would still leave the height limited by the shorter line. 
	// Moving the pointer with the shorter line increases the likelihood of finding a taller line that can result in a larger area, even though the width decreases.
	for s < e {
		w := e - s
		h := nums[s]
		if nums[s] >= nums[e] {
			h = nums[e]
			e--
		} else {
			s++
		}
		tAns := w * h
		if tAns > ans {
			ans = tAns
		}
	}
	return ans
}
