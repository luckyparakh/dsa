package main

func productExceptSelf(nums []int) []int {
	// TC:=o(n^2)
	l := len(nums)
	ans := make([]int, l)
	for i := range l {
		m := 1
		for j := range l {
			if i != j {
				m *= nums[j]
			}
		}
		ans[i] = m
	}
	return ans
}

func productExceptSelf2(nums []int) []int {
	// TC:=o(n)
	l := len(nums)
	ans := make([]int, l)
	prefix := make([]int, l)
	prefix[0] = 1
	suffix := make([]int, l)
	suffix[l-1] = 1
	for i := 1; i < l; i++ {
		// (multiplication till now) * (multiple previous element)
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := l-2; i >=0; i-- {
		// (multiplication till now) * (multiple after element)
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i:=range l{
		ans[i]=prefix[i]*suffix[i]
	}
	return ans
}
