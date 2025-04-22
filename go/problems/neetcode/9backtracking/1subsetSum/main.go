package main

func subsets(nums []int) [][]int {
	var h func([]int, int)
	ans := [][]int{}
	ans = append(ans, []int{}, nums)
	h = func(arr []int, start int) {
		l := len(arr)
		if start < l {
			ans = append(ans, []int{arr[start]})
		}
		for i := start + 1; i < l; i++ {
			
		}
	}
	h(nums, 0)
	return ans
}
