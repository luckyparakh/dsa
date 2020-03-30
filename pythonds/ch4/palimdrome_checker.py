# -*- coding: utf-8 -*-

from ch4.deque import deque

def palindrome_check(palin):
    d=deque()
    flag=True
    for i in palin:
        d.addFront(i)  
    while(d.size() > 1):
        if(d.removeFront() != d.removeRear()):
            flag=False
            break
    return flag
print(palindrome_check("radar"))