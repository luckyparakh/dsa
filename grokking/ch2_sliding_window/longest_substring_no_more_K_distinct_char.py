# -*- coding: utf-8 -*-

'''
Given a string, find the length of the longest substring in it with no more 
than K distinct characters.
LeetCode:
TC:
SC:
'''
def long_substr_with_k_distinct(word,k):
    temp_st=''
    start=0
    max_len=0
    distinct_letter=0
    for end in range(len(word)):
        temp_st+=word[end]
        distinct_letter=len(set(temp_st))
        
        '''
        In case, need to print string with K distinct letters
        if (distinct_letter==k):
            print(temp_st)'''
        while (distinct_letter>k):
            start+=1
            temp_st=temp_st[1:]
            distinct_letter=len(set(temp_st))
            '''
            In case, need to print string with K distinct letters
            if (distinct_letter==k):
                print(temp_st)'''
        max_len=max(max_len,end-start+1)
    return 0 if max_len==0 else max_len

print(long_substr_with_k_distinct("araaci",2))         
print(long_substr_with_k_distinct("araaci",1))
print(long_substr_with_k_distinct("cbbebi",3))
print(long_substr_with_k_distinct("aabbcc",1))
print(long_substr_with_k_distinct("aabbcc",2))   
print(long_substr_with_k_distinct("aabbcc",3))
print(long_substr_with_k_distinct("aaabbb",3)) 
print(long_substr_with_k_distinct("aabbaacdeeeeddded",3)) 
print(long_substr_with_k_distinct("acddefabc",4)) 
print(long_substr_with_k_distinct("aaaabbbb",4)) 