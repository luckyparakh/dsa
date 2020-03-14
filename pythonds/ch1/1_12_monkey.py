# -*- coding: utf-8 -*-
"""
https://runestone.academy/runestone/books/published/pythonds/Introduction/DefiningFunctions.html
The theorem states that a monkey hitting keys at random on a typewriter keyboard for an infinite amount of time 
will almost surely type a given text, such as the complete works of William Shakespeare. 
Well, suppose we replace a monkey with a Python function. 
How long do you think it would take for a Python function to generate just one sentence of Shakespeare? 
The sentence we’ll shoot for is: “methinks it is like a weasel”

The way we’ll simulate this is to write a function that generates a string that is 28 characters long by 
choosing random letters from the 26 letters in the alphabet plus the space. 
We’ll write another function that will score each generated string by comparing the randomly generated string to the goal.

A third function will repeatedly call generate and score, then if 100% of the letters are correct we are done. 
If the letters are not correct then we will generate a whole new string.To make it easier to follow your program’s 
progress this third function should print out the best string generated so far and its score every 1000 tries.
"""
import random
word_list=[" "]
goal_sentence="methinks it is like a weasel"

def generate_word_list():
    for i in range(97,123):
        word_list.append(chr(i))
    print("Word List is equal to:{}".format(word_list))

def random_generator():
    s=""
    for i in range(28):
        s=s+random.choice(word_list)
    return s

def score_generator(goal,test): 
    score=0
    for i in range(len(goal)):
        if(goal[i]==test[i]):
            score=score+1
    score=score/len(goal)
    return score

def main():
    i=0
    generate_word_list()
    random_sentence=random_generator()
    score=score_generator(goal_sentence,random_sentence)
    while score < 1:
        if(i%10000==0):
            print("Random_sentence generated at {} is equal to:{} & score is {}.".format(i,random_sentence,score))
        random_sentence=random_generator()
        score=score_generator(goal_sentence,random_sentence)
        i=i+1

main()