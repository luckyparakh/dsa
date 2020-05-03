# -*- coding: utf-8 -*-

'''
Given a string and a pattern, find all anagrams of the pattern in the given string.
LeetCode:438 anautation in String
TC: O(N)
SC: O(M+M) where M is lenght of ana_st
'''

def anagram_in_str(st,ana_st):
    start=0
    ana_freq={}
    char_freq={}
    ana_start=[]
    for letter in ana_st:
        if letter not in ana_freq:
            ana_freq[letter]=0
        ana_freq[letter]+=1
    for end,letter in enumerate(st):
        if letter not in char_freq:
            char_freq[letter]=0
        char_freq[letter]+=1
        if end >= len(ana_st)-1:
            if char_freq==ana_freq:
                ana_start.append(start)
            start_letter=st[start]
            char_freq[start_letter]-=1
            if char_freq[start_letter]==0:
                del char_freq[start_letter]
            start+=1
    return ana_start

print(anagram_in_str("ppqp","pq"))
print(anagram_in_str("abbcabc","abc"))
print(anagram_in_str("cbaebabacd" , "abc"))
print(anagram_in_str("abab","ab"))