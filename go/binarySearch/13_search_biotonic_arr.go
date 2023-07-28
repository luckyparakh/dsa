package main

// https://leetcode.com/problems/find-in-mountain-array/
import "fmt"

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	array []int
}

func (this *MountainArray) get(index int) int {
	return this.array[index]
}
func (this *MountainArray) length() int {
	return len(this.array)
}
func findPeakElement(nums *MountainArray) int {
	l := nums.length()
	s := 0
	e := l - 1
	for s <= e {
		m := s + (e-s)/2
		val := nums.get(m)
		if m == 0 {
			if val > nums.get(m+1) {
				return m
			}
			return m + 1
		}
		if m == l-1 {
			if val > nums.get(m-1) {
				return m
			}
			return m - 1
		}
		val1 := nums.get(m + 1)
		if val > nums.get(m-1) && val > val1 {
			return m
		} else if val < val1 {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return -1
}
func bs(n *MountainArray, s, e, t int, asc bool) int {
	for s <= e {
		m := s + (e-s)/2
		val := n.get(m)
		if val == t {
			return m
		}
		if val > t {
			if asc {
				e = m - 1
			} else {
				s = m + 1
			}

		} else {
			if asc {
				s = m + 1
			} else {
				e = m - 1
			}
		}
	}
	return -1
}
func findInMountainArray(target int, mountainArr *MountainArray) int {
	peak := findPeakElement(mountainArr)
	op := bs(mountainArr, 0, peak, target, true)
	if op == -1 {
		return bs(mountainArr, peak+1, mountainArr.length()-1, target, false)
	}
	return op
}

func main() {
	ma := MountainArray{
		array: []int{1, 2, 3, 4, 5, 3, 1},
	}
	fmt.Println(findInMountainArray(3, &ma))
	ma1 := MountainArray{
		array: []int{1, 2, 3, 4, 5, 3, 1},
	}
	fmt.Println(findInMountainArray(7, &ma1))
}
