# -*- coding: utf-8 -*-

'''
Given a string, find the length of the longest substring which has no repeating characters.
LeetCode:3 Longest Substring Without Repeating Characters
TC:O(n)
SC:O(n) due to dictionary used

'''

def no_repeat_substring(word):
    start=end=max_len=0
    letter_index={}
    end_m=start_m=0
    for end,letter in enumerate(word):
        if letter in letter_index:
            start=max(start,letter_index[letter]+1)
        letter_index[letter]=end
        '''
        Save value of end & start when lenght of no repeat SubStr is max.
        if end-start+1 > max_len:
            end_m=end
            start_m=start
        '''
        max_len=max(max_len,end-start+1)
        
    '''Print max no repeat SubStr'''
    #print(word[start_m:end_m+1])
    return max_len
print(no_repeat_substring("aabccbb"))
print(no_repeat_substring("abbbed"))
print(no_repeat_substring("abcdrstcfba"))
print(no_repeat_substring("bbbb"))
print(no_repeat_substring("abcabcbb"))
print(no_repeat_substring("pwwkew"))
print(no_repeat_substring("dvdf"))