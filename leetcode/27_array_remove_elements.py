# -*- coding: utf-8 -*-
def removeElement(nums,val):
    '''
    Time = O(n)
    Space = O(1)
    '''
    j=0
    for i in range(0,len(nums)):
        if(nums[i]!=val):
            nums[j]=nums[i]
            j=j+1
    return nums[:j]
num=[0,1,0,4,12]  
print(removeElement(num,1))