# -*- coding: utf-8 -*-

'''
Given an array of positive numbers and a positive number ‘k’, 
find the maximum sum of any contiguous subarray of size ‘k’.
'''

def max_sum_cont_subarray(nums,k):
    start=0
    win_sum=0
    sum_list=[]
    for win_end in range(len(nums)):
        win_sum+=nums[win_end]
        if win_end>=k-1:
            sum_list.append(win_sum)
            win_sum-=nums[start]
            start+=1
    return max(sum_list)

print(max_sum_cont_subarray([2, 1, 5, 1, 3, 2],3))
print(max_sum_cont_subarray([2, 3, 4, 1, 5],2))