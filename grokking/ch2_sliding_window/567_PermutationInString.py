# -*- coding: utf-8 -*-

'''
Given a string and a pattern, find out if the string contains any permutation of the pattern.
LeetCode:567. Permutation in String
'''
'''
TC: O(N*N)
SC: O(1)
'''

def permutation_in_str(perm_st,st):
    for start,letter in enumerate(st):
        end=start
        if letter in perm_st:
            temp_st=perm_st
            while(end<len(st) and (end-start+1)<=len(perm_st)):
                if st[end] in temp_st:
                    pos=temp_st.find(st[end])
                    temp_st=temp_st[:pos]+temp_st[pos+1:]
                else:
                    break
                if temp_st=="":
                    return True
                end+=1
    return False

'''
TC: O(N)
SC: O(M+M) where M is lenght of perm_st
'''
def permutation_in_str_eff(perm_st,st):
    start=0
    perm_freq={}
    char_freq={}
    for letter in perm_st:
        if letter not in perm_freq:
            perm_freq[letter]=0
        perm_freq[letter]+=1
    for end,letter in enumerate(st):
        if letter not in char_freq:
            char_freq[letter]=0
        char_freq[letter]+=1
        if end >= len(perm_st)-1:
            if char_freq==perm_freq:
                return True
            start_letter=st[start]
            char_freq[start_letter]-=1
            if char_freq[start_letter]==0:
                del char_freq[start_letter]
            start+=1
    return False
            
    

print(permutation_in_str("ab","eidbaooo"))
print(permutation_in_str("ab","eidboaoo"))
print(permutation_in_str("abc","oidbcaf"))
print(permutation_in_str("dc","odicf"))
print(permutation_in_str("bcdxabcdy","bcdyabcdx"))
print(permutation_in_str("aabc","aaacb"))
print(permutation_in_str("abc","babaabbcb"))
print(permutation_in_str("adc","dadc"))

print("===============================")
print(permutation_in_str_eff("ab","eidbaooo"))
print(permutation_in_str_eff("ab","eidboaoo"))
print(permutation_in_str_eff("abc","oidbcaf"))
print(permutation_in_str_eff("dc","odicf"))
print(permutation_in_str_eff("bcdxabcdy","bcdyabcdx"))
print(permutation_in_str_eff("aabc","aaacb"))
print(permutation_in_str_eff("abc","babaabbcb"))
print(permutation_in_str_eff("adc","dadc"))