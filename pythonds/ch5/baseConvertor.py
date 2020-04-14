# -*- coding: utf-8 -*-

#Digit Convertor via Recursion"

def toStr(n,base):
    digit="0123456789abcdef"
    if n < base:
        return digit[n]
    else:
        return toStr(n//base,base) + digit[n%base]

print(toStr(1453,16))
print(toStr(756,10))
print(toStr(10,2))