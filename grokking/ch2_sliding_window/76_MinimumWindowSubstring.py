# -*- coding: utf-8 -*-

'''
Given a string S and a string T, find the minimum window in S 
which will contain all the characters in T in complexity O(n).
LeetCode:76. Minimum Window Substring
TC: O(N+M)
SC: O(N) where n is size of pattern
'''
def minWindow(st,pattern):
    start=match=0
    min_len=len(st)+1
    pattern_freq={}
    for letter in pattern:
        if letter not in pattern_freq:
            pattern_freq[letter]=0
        pattern_freq[letter]+=1
    for end,letter in enumerate(st):
        if letter in pattern_freq:
            pattern_freq[letter]-=1
            if pattern_freq[letter]>=0:
               match+=1
        
        while (match==len(pattern)):
            if min_len>end-start+1:
                min_len=end-start+1
                substr=st[start:end+1]
            left_letter=st[start]
            if left_letter in pattern:
                if pattern_freq[left_letter]>=0:
                    match-=1
                pattern_freq[left_letter]+=1
            start+=1
       
    return "" if min_len>len(st) else substr

print(minWindow("AADOBECODEBANC","AABC"))
print(minWindow("aabdec","abc"))
print(minWindow("ababdaabca","aabc"))
print(minWindow("abeeedc","abc"))
print(minWindow("eaeabc","aab"))
print(minWindow("adcad","abc"))
print(minWindow("bdab","ab"))