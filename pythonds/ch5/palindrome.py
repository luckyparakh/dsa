# -*- coding: utf-8 -*-

#palindrome via recursion

def isPalin(s):
    if len(s) == 2:
        if s[0] == s[-1]:
            return True
        else:
            return False
    elif len(s)==1:
        return True
    else:
        if s[0] == s[-1]:
            return isPalin(s[1:-1])
        else:
            return False

def generate_word_list():
    word_list=""
    for i in range(97,123):
        word_list=word_list+(chr(i))
    return word_list
    
def removeOtherChar(s):
    stringWithNoSpecialChar=""
    s=s.lower()
    for i in s:
        if i in letter_list:
            stringWithNoSpecialChar=stringWithNoSpecialChar+i
    return stringWithNoSpecialChar
    
letter_list=generate_word_list()
#print(letter_list)

print(isPalin(removeOtherChar("kayak")))
print(isPalin(removeOtherChar("aibohphobia")))
print(isPalin(removeOtherChar("Live not on evil")))
print(isPalin(removeOtherChar("Reviled did I live, said I, as evil I did deliver")))
print(isPalin(removeOtherChar("Go hang a salami; I’m a lasagna hog.")))