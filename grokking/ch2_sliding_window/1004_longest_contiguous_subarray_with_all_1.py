# -*- coding: utf-8 -*-

'''
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s,
find the length of the longest contiguous subarray having all 1s.
LeetCode:1004 Max Consecutive Ones III

TC:O(N+N)
'''

def findlongestConsecutiveOnes(nums, k) :
    start=max_len=count_zero=0
    for end,num in enumerate(nums):
        if (num==0):
            count_zero+=1
        while (count_zero>k):
            if(nums[start]==0):
                count_zero-=1
            start+=1
        max_len=max(max_len,end-start+1)
    return max_len

print(findlongestConsecutiveOnes([1,1,1,0,0,0,1,1,1,1,0],2))
print(findlongestConsecutiveOnes([0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1],3))
print(findlongestConsecutiveOnes([0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1],2))
print(findlongestConsecutiveOnes([0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1],3))