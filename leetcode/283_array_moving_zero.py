from datetime import datetime

def moveZeroes(nums):
        """
        Time=O(n2)
        Space=O(1)
        """
        no_zero=nums.count(0)
        for i in range(0,no_zero):
            for j in range(0,len(nums)-1):
                if(nums[j]==0):
                    nums[j]=nums[j+1]
                    nums[j+1]=0
        print(nums)
def moveZeroes1(nums):
        """
        Do not return anything, modify nums in-place instead.
        """
        j=0
        for i in range(0,len(nums)):
            if(nums[i]!=0):
                nums[j]=nums[i]
                j=j+1
        while (j<len(nums)):
            nums[j]=0
            j=j+1
                
                
        print(nums)
nums=[0,1,0,4,12]  
start=datetime.now()
#moveZeroes(nums)
print("time:",datetime.now()-start)
start=datetime.now()
moveZeroes1(nums)
print("time:",datetime.now()-start)


