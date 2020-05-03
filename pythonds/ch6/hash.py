# -*- coding: utf-8 -*-

'''
You can probably already see that this technique is going to work only if each item maps to a unique location in the hash table. 
For example, if the item 44 had been the next item in our collection, it would have a hash value of 0 (44%11==0). 
Since 77 also had a hash value of 0, we would have a problem. According to the hash function, two or more items would need to be in the same slot. 
This is referred to as a collision (it may also be called a “clash”). 
'''
def remainder_hash(n):
    return n%size_hash
def store_hash(alist):
    for i in alist:
        hash_val=remainder_hash(i)
        hash_table[hash_val]=i


size_hash=11
hash_table=[None]*size_hash
alist=[54,26,93,17,77,31]
store_hash(alist)
print(hash_table)