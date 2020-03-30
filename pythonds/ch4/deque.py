# -*- coding: utf-8 -*-
class deque:
    #Considering 0th element as rear
    def __init__(self):
        self.item=[]
    def isEmpty(self):
        return self.item==[]
    def size(self):
        return len(self.item)
    def addFront(self,val):
        self.item.append(val)
    def addRear(self,val):
        self.item.insert(0,val)
    def removeFront(self):
        return self.item.pop()
    def removeRear(self):
        return self.item.pop(0)

if (__name__=='__main__'):
    d=deque()
    print(d.isEmpty())
    d.addRear(4)
    d.addRear(40)
    d.addFront("dog")
    d.addRear(8.4)
    d.addFront("cat")
    print(d.item)
    print(d.size())
    print(d.isEmpty())
    d.removeFront()
    d.removeRear()
    print(d.item)