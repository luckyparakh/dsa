# -*- coding: utf-8 -*-
'''
Given an array of positive numbers and a positive number ‘k’, 
find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘k’. 
Return 0, if no such subarray exists.

LeetCode
209. Minimum Size Subarray Sum

TC: O(N+N) ( A N for 'for' loop & another N for while loop)
SC: O(1)
'''

def smallest_sub_array(ll,k):
    sum_ele=start=0
    min_arr_len=float('inf')
    for end in range(len(ll)):
        sum_ele+=ll[end]
        while (sum_ele>=k):
            min_arr_len=min(min_arr_len,end-start+1)
            sum_ele-=ll[start]
            start+=1
    return 0 if min_arr_len==float('inf') else min_arr_len
        
    
print(smallest_sub_array([2, 1, 5, 2, 3, 2], 7 ))
print(smallest_sub_array([2, 1, 5, 2, 8], 7 ))
print(smallest_sub_array([3, 4, 1, 1, 6], 100)) 
print(smallest_sub_array([1,2,3,4,5], 15)) 