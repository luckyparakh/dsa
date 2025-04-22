package main

func minCostClimbingStairs(cost []int) int {
    size := 1001
	memo := make([]int, size)
	for i := 0; i < size; i++ {
		memo[i] = -1
	}
	var helper func(int, []int) int
	helper = func(index int, memo []int) int {
        if index==-1{
            return 0
        }
		if memo[index] != -1 {
			return memo[index]
		}
// Cost is zero, as from step zero and 1, cost is zero
		if index == 0 || index == 1 {
			return 0
		}
        sum:=min((helper(index-1, memo)+cost[index-1]), (helper(index-2, memo)+cost[index-2]))
		memo[index]=sum
        return sum
	}
	return helper(len(cost), memo)
}