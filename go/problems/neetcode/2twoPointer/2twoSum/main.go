package main

func twoSum(numbers []int, target int) []int {
	s, e := 0, len(numbers)-1
	ans := []int{}
	for s < e {
		currentSum := numbers[s] + numbers[e]
		if currentSum == target {
			ans = append(ans, s+1, e+1)
			break
		}
		if currentSum > target {
			e--
		}
		if currentSum < target {
			s++
		}
	}
	return ans
}
