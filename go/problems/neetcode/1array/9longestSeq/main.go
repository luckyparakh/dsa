package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	//tle
	// because it has 2 loops
	hash := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = true
		}
	}
	fmt.Println(hash)
	ans := 0
	for i := range nums {
		l := 1
		t := nums[i]
		for {
			t += 1
			if _, ok := hash[t]; ok {
				l++
			} else {
				break
			}
		}
		ans = max(ans, l)
	}
	return ans
}
func longestConsecutive1(nums []int) int {
	//tle
	if len(nums) == 0 {
		return 0
	}
	hash := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = true
		}
	}
	// fmt.Println(hash)
	ans := 0
	for i := range nums {
		t := nums[i]
		// By adding this condition, it is ensured that it is start of sequence, but this has also TLE because we iterating over nums
		// And nums can have duplictae elements like [1,1,0,2,3], here 1 is duplictae and this login will twice to calculate the same ans
		if _, ok := hash[t-1]; !ok {
			l := 1
			for {
				t += 1
				if _, ok := hash[t]; ok {
					l++
				} else {
					break
				}
			}
			if l > ans {
				ans = l
			}
		}
	}
	return ans
}
func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hash := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = true
		}
	}
	// fmt.Println(hash)
	ans := 0
	// By adding this condition, it is ensured that it is start of sequence, but this has also TLE because we iterating over nums
		// And nums can have duplictae elements like [1,1,0,2,3], here 1 is duplictae and this login will twice to calculate the same ans
		// iterate over hash is good way and by this way a number is only accessed twice.
	for i := range hash {
		if _, ok := hash[i-1]; !ok {
			l := 1
			for {
				i += 1
				if _, ok := hash[i]; ok {
					l++
				} else {
					break
				}
			}
			if l > ans {
				ans = l
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
