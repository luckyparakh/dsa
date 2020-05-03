# -*- coding: utf-8 -*-
"""
Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
LeetCode: 643. Maximum Average Subarray I
"""
def cont_subarr_avg(ll,k):
    """
    Brute Force
    O(N∗K) where ‘N’ is the number of elements in the input list.
    The inefficiency is that for any two consecutive subarrays of size ‘K’, 
    the overlapping part (which will contain k-1 elements) will be evaluated twice.
    """
    start=0
    end=k
    avg_list=[]
    while(end<=len(ll)):
        sum=0
        for i in range(start,end):
            sum=sum+ll[i]
        start+=1
        end+=1
        avg_list.append(sum/k)
    return avg_list
    #return max(avg_list)

def cont_subarr_avg_sw(ll,k):
    """
    Sliding Window
    O(N) where ‘N’ is the number of elements in the input list.
    """
    start=0
    avg_list=[]
    sum=0
    for end in range(len(ll)):
        sum+=ll[end]
        #Slide the window, only in case when window size of k is hit.
        if(end>=k-1):
            avg_list.append(sum/k)
            sum-=ll[start]
            start+=1
    return avg_list
    #return max(avg_list)

print(cont_subarr_avg([1, 3, 2, 6, -1, 4, 1, 8, 20],5))
print(cont_subarr_avg_sw([1, 3, 2, 6, -1, 4, 1, 8, 20],5))