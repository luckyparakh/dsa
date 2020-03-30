# -*- coding: utf-8 -*-

class queue:
    
    def __init__(self):
        self.item=[]
    def isEmpty(self):
        return self.item==[]
    def size(self):
        return len(self.item)
    def enqueue(self,val):
        self.item.insert(0,val) #O(N)
    def dequeue(self):
        return self.item.pop() #O(1)

if (__name__=='__main__'):      
    q=queue()
    print(q.isEmpty())
    q.enqueue(14)
    q.enqueue('dog')
    q.enqueue(True)
    q.enqueue(8.4)
    print(q.item)
    print(q.size())
    print(q.dequeue())
    print(q.isEmpty())
    print(q.dequeue())
    print(q.item[0])