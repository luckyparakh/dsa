# -*- coding: utf-8 -*-

def binarySearch(numbers,n):
    '''
    O(n)=logn
    '''
    start=0
    end=len(numbers)
    found=False
    while (start<=end):
        mid=(start+end)//2
        if numbers[mid]==n:
            found=True
            break
        else:
            if numbers[mid]<n:
                start=mid+1
            else:
                end=mid-1
    return found

def binarySearchRecur(alist,item):
    if len(alist)==0:
        return False
    else:
        mid=len(alist)//2
        if alist[mid]==item:
            return True
        else:
            if alist[mid]>item:
                return binarySearchRecur(alist[:mid],item)
            else:
                return binarySearchRecur(alist[mid+1:],item)

testlist = [0, 1, 2, 8, 13, 17, 19, 32, 42,]
print(binarySearch(testlist, 3))
print(binarySearch(testlist, 13))
print("binarySearchRecur")
print(binarySearchRecur(testlist, 3))
print(binarySearchRecur(testlist, 13))