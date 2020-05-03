# -*- coding: utf-8 -*-


def validMountain(nums):
    '''
    Time = 
    Space = 
    '''
    if(len(nums)>=3):
        isValid=True
        
        
        if(nums[0]<nums[1]):
            prev_slope=1
        elif(nums[0]>nums[1]):
            prev_slope=-1
        else:
            prev_slope=0
        
        if(prev_slope==0):
            isValid=False
            return False
        peak=0
        current_slope=0
        for i in range(1,len(nums)-1):
            
            if(nums[i]<nums[i+1]):
                current_slope=1
            elif(nums[i]>nums[i+1]):
                current_slope=-1
            else:
                current_slope=0
            
            if(current_slope==0):
                isValid=False
                break
            if(current_slope==-1 and prev_slope==1):
                peak=peak+1
            if(current_slope==1 and prev_slope==-1):
                isValid=False
                break
            prev_slope=current_slope
        if(peak==0):
            isValid=False
        return isValid
    else:
        return False
num=[0,3,4,12,1,0] 
print(validMountain(num))
num=[1,3,2]  
print(validMountain(num))
num=[0,1,0,4,12]  
print(validMountain(num))
num=[0,0,4,12,1] 
print(validMountain(num))
num=[12,1]
print(validMountain(num))
num=[0,3,2,1] 
print(validMountain(num))
