# -*- coding: utf-8 -*-
class Node:
    def __init__(self,initval):
        self.data=initval
        self.next=None
    def getData(self):
        return self.data
    def setData(self,newdata):
        self.data=newdata
    def getNext(self):
        return self.next
    def setNext(self,newnext):
        self.next=newnext

class unOrderedList:
    def __init__(self):
        self.head=None
    def isEmpty(self):
        return self.head == None
    def add(self,val):
        temp=Node(val)
        temp.setNext(self.head)
        self.head=temp
    def size(self):
        current=self.head
        count=0
        while (current!=None):
            count=count+1
            current=current.getNext()
        return count
    def search(self,val):
        current=self.head
        found=False
        while (current!=None):
            if(current.getData()==val):
                found=True
                break
            current=current.getNext()
        return found
    def remove(self,val):
        current=self.head
        previous=None
        while (current!=None):
            if(current.getData()==val):
                break
            previous=current
            current=current.getNext()
        if previous==None:
            self.head=current.getNext()
        elif current==None:
            raise ValueError("Element not found")
        else:
            previous.setNext(current.getNext())
    def pop(self):
        value=self.head.getData()
        self.head=self.head.getNext()
        return value
    def index(self,val):
       pos=1
       current=self.head
       while(current!=None):
           if(current.getData()==val):
               break
           pos=pos+1
           current=current.getNext()
       if current!=None:
           return pos
       else:
           raise ValueError("Value is not present")
    def traverse(self):
        current=self.head
        if(self.isEmpty()):
            print("Empty List")
        else:
            while (current!=None):
                print(current.getData(),end=" ")
                current=current.getNext()
            print()

if __name__=='__main__':
    ull=unOrderedList()
    ull.add(10)
    ull.remove(10)
    ull.traverse()
    ull.add(20)
    ull.add(30)
    print("size",ull.size())
    ull.traverse()
    print("poped:",ull.pop())
    ull.traverse()
    print("search",ull.search(20))
    print("search",ull.search(30))
    ull.add(40)
    ull.add(50)
    ull.add(60)
    ull.traverse()
    ull.remove(50)
    ull.traverse()
    ull.remove(60)
    ull.traverse()
    print("Index",ull.index(40))
    ull.remove(20)
    ull.traverse()