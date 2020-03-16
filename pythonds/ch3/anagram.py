# -*- coding: utf-8 -*-
'''
One string is an anagram of another if the second is simply a rearrangement of the first. For example, 'heart' and 'earth' are anagrams.
The strings 'python' and 'typhon' are anagrams as well. 
For the sake of simplicity, we will assume that the two strings in question are of equal length and that they are made up of symbols from the set of 26 lowercase alphabetic characters. 
Our goal is to write a boolean function that will take two strings and return whether they are anagrams
'''
import time
def anagram(s1,s2):
    '''
    TC=n*(n+1)/2
    O(n2)
    '''
    flag=True
    if (len(s1)==len(s2)):
        s1_list=list(s1)
        for char in s2:
            if char in s1_list:
                char_pos=s1_list.index(char)
                s1_list[char_pos]=None
            else:
                flag=False
                break
    else:
        flag=False
    return flag
def anagram_sort(s1,s2):
    '''
    Time Complexity depends on sort function.
    Typically sort function's TC is O(n2) or O(nlogn)
    '''
    flag=True
    if (len(s1)==len(s2)):
        s1_list=list(s1)
        s2_list=list(s2)
        s1_list.sort()
        s2_list.sort()
        for i in range(len(s1_list)):
            if(s1_list[i]!=s2_list[i]):
                flag=False
                break
    else:
        flag=False
    return flag
def anagram_strike_off(s1,s2):
    '''    
        O(n)
    '''
    flag=True
    c=[0]*26
    pos=0
    if (len(s1)==len(s2)):
        for char in s1:
            pos=ord(char)-ord('a')
            c[pos]=1
        for char in s1:
            pos=ord(char)-ord('a')
            if(c[pos]==1):
                c[pos]=-1
            else:
                flag=False
                break
        for i in c:
            if(c[i]!=-1):
                flag=False
                break
    else:
        flag=False
    return flag
    
s1="earth"
s2="heatt"
start_time=time.time()
op=anagram(s1,s2)
end_time=time.time()
print("anagram_output:",op)
op=anagram_sort(s1,s2)
print("anagram_sort_output:",op)
op=anagram_strike_off(s1,s2)
print("anagram_sort_output:",op)