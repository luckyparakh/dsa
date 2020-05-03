# -*- coding: utf-8 -*-

'''
Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’ letters with any letter,
find the length of the longest substring having the same letters after replacement.
LeetCode
TC
SC
'''

def longest_substr_same_letter_replace(word,k):
    start=end=max_len=0
    
    for start,letter in enumerate(word):
        end=start
        max_replace=k
        
        while (max_replace>0 and end < len(word)):
            if word[start]!=word[end]:
                max_replace-=1
            end+=1
            
            '''
            if max_replace>0:
                if word[start]!=word[end]:
                    max_replace-=1
                end+=1
            else:
                if word[start]==word[end]:
                    end+=1
                else:
                    break
            '''
        max_len=max(max_len,end-start+1)
    return max_len

print(longest_substr_same_letter_replace("aabccbbe",2))
print(longest_substr_same_letter_replace("abbcb",1))
print(longest_substr_same_letter_replace("abccc",2))
