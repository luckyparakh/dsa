# -*- coding: utf-8 -*-

'''
Given a string and a list of words, find all the starting indices of substrings in the given string that are a concatenation of all the given words exactly once without any overlapping of words. 
It is given that all words are of the same length.
LeetCode:30
TC
SC
'''
'''
def findSubstring(st,words):
    word_freq={}
    match=start=end=0
    substr_index=[]
    word_len=len(words[0])
    win_size=len(words)*len(words[0])
    for word in words:
        if word not in word_freq:
            word_freq[word]=0
        word_freq[word]+=1
    #while end<len(st):
    for end in range(len(st)):
        word=st[end+1-word_len:end+1]
        #word=st[end:end+word_len]
        if word in word_freq:
            word_freq[word]-=1
            if word_freq[word]>=0:
                match+=1
            
        if match==len(words):
            substr_index.append(start)
        if end-start+1==win_size:
            left_word=st[start:start+word_len]
            if left_word in word_freq:
                if word_freq[left_word]>=0:
                    match-=1
                word_freq[left_word]+=1
                start=start+(word_len-1)
            else:
                start+=1
        #end+=1
    return substr_index
'''
def findSubstring(st,words):
    word_freq={}
    substr_index=[]
    word_len=len(words[0])
    win_size=len(words)*len(words[0])
    for word in words:
        if word not in word_freq:
            word_freq[word]=0
        word_freq[word]+=1
    for i in range(len(st)-win_size+1):
        #words_seen={}
        word_freq_copy=word_freq.copy()
        for j in range(len(words)):
            word=st[i+j*word_len:i+(j+1)*word_len]
            if word not in word_freq:
                break
            else:
                word_freq_copy[word]-=1
                if word_freq_copy[word]<0:
                    break
            if j+1==len(words):
                substr_index.append(i) 
                
            
    return substr_index
            
print(findSubstring("barfoothethefoobarman",["foo","bar"]))
print(findSubstring("wordgoodgoodgoodbestword", ["word","good","best","word"]))
print(findSubstring("catfoxcat",["cat", "fox"]))
print(findSubstring("catcatfoxfox",["cat", "fox"]))
print(findSubstring("barfoofoobarthefoobarman",["bar","foo","the"]))
print(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake",["fooo","barr","wing","ding","wing"]))
print(findSubstring("a",["a"]))
print(findSubstring("aaaaaaaa",["aa","aa","aa"]))
print(findSubstring("aaaaaaaa",["aa","aa","aa","aa"]))