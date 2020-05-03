# -*- coding: utf-8 -*-
'''

Given an array of characters where each character represents a fruit tree, you are given two baskets and your goal is to put maximum number of fruits in each basket.
The only restriction is that each basket can have only one type of fruit.
You can start with any tree, but once you have started you can’t skip a tree.
You will pick one fruit from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.
Write a function to return the maximum number of fruits in both the baskets.

LeetCode:904. Fruit Into Baskets
TC:O(N+N)
SC:O(N)
'''

def fruits_in_basket(trees):
    start=end=max_fruit=0
    freq_tree={}
    for end,tree in enumerate(trees):
        if tree not in freq_tree:
            freq_tree[tree]=0
        freq_tree[tree]=freq_tree[tree]+1
        while len(freq_tree)>2:
            freq_tree[trees[start]]-=1
            if freq_tree[trees[start]]==0:
                del freq_tree[trees[start]]
            start+=1
        max_fruit=max(max_fruit,end-start+1)
    return max_fruit

print(fruits_in_basket(['A', 'B', 'C', 'B', 'B', 'C']))
print(fruits_in_basket(['A', 'B', 'C', 'A', 'C']))
print(fruits_in_basket([1,2,1]))
print(fruits_in_basket([0,1,2,2]))
print(fruits_in_basket([1,2,3,2,2]))
print(fruits_in_basket([3,3,3,1,2,1,1,2,3,3,4]))