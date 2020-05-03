# -*- coding: utf-8 -*-
'''
With DP

A classic example of an optimization problem involves making change using the fewest coins. 
Suppose you are a programmer for a vending machine manufacturer. Your company wants to streamline effort by giving out the fewest possible coins in change for each transaction. 
Suppose a customer puts in a dollar bill and purchases an item for 37 cents. 
What is the smallest number of coins you can use to make change? The answer is six coins: two quarters, one dime, and three pennies. 
How did we arrive at the answer of six coins? 
We start with the largest coin in our arsenal (a quarter) and use as many of those as possible, then we go to the next lowest coin value and use as many of those as possible. 
This first approach is called a greedy method because we try to solve as big a piece of the problem as possible right away.

The greedy method works fine when we are using U.S. coins, but suppose that your company decides to deploy its vending machines in Lower Elbonia where, in addition to the usual 1, 5, 10, and 25 cent coins they also have a 21 cent coin. In this instance our greedy method fails to find the optimal solution for 63 cents in change. With the addition of the 21 cent coin the greedy method would still find the solution to be six coins. However, the optimal answer is three 21 cent pieces.
'''
def recMC(coinValueList,change):
    minCoins=change/min(coinValueList)
    if change in coinValueList:
        return 1
    else:
        if change in coinChange.keys():
            return coinChange[change]
        else:
            for i in coinValueList:
                if i <= change:
                    numCoins=1+recMC(coinValueList,change-i)
                    if numCoins<minCoins:
                        minCoins=numCoins
            coinChange[change]=minCoins
            return coinChange[change]

coinChange={}
print(recMC([1,5,10,21,25],63))