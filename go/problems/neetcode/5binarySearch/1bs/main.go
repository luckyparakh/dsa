package main

func search(nums []int, target int) int {
	s, e := 0, len(nums)-1

	for s <= e {
		m := s + (e-s)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			e = m-1
		} else {
			s = m+1
		}
	}
	return -1
}
