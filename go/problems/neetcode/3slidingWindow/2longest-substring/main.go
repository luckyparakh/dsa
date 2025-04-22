package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	d := make(map[byte]int)
	ans := 0
	start := 0
	for end := range s {
		// fmt.Println("End", end)
		// Below condition tries to manage windows length
		if _, ok := d[s[end]]; ok {
			// fmt.Println("In ok", d)
			till:=d[s[end]]
			for i := start; i <=till; i++ {
				delete(d, s[i])
			}
			start = till + 1
			// fmt.Println("In ok after for loop", d, start)
		}
		// To calculate the answer
		if end-start+1 > ans {
			ans = end - start + 1
		}
		// fmt.Println("ans", ans)
		d[s[end]] = end
	}
	return ans
}

func main() {
	// fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	// fmt.Println(subarraySumFixed([]int{1, 2, 3, 7, 4, 1}, 3))
}

func subarraySumFixed(nums []int, k int) int {
	s, ans, start := 0, 0, 0
	for i := range nums {
		if i-start >= k {
			s -= nums[start]
			start++
		}
		s += nums[i]
		if s > ans {
			ans = s
		}
	}
	return ans
}

func subarraySumLongest(nums []int, k int) int {
	tempSum, ans, start := 0, 0, 0
	for end := range nums {
		// Shrink window
		for tempSum > k {
			tempSum -= nums[start]
			start++
		}
		l := end - start + 1
		tempSum += nums[end]
		if l > ans {
			ans = l
		}
	}
	return ans
}
