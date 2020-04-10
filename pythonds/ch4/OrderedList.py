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

class OrderedList:
    def __init__(self):
        self.head=None
    def isEmpty(self):
        return self.head == None
    def size(self):
        current=self.head
        count=0
        while (current!=None):
            count=count+1
            current=current.getNext()
        return count
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
            print("Element not found")
        else:
            previous.setNext(current.getNext())
    def search(self,val):
        current=self.head
        found=False
        while (current!=None):
            if(current.getData()==val):
                found=True
                break
            if(current.getData()>val):
                break
            current=current.getNext()
        return found   
    def pop(self):
        value=self.head.getData()
        self.head=self.head.getNext()
        return value
    def add(self,val):
        previous=None
        current=self.head
        temp=Node(val)
        while(current!=None):
            if(current.getData()>val):
                break
            previous=current
            current=current.getNext()
        if previous==None:
            temp.setNext(self.head)
            self.head=temp
        else:
            temp.setNext(previous.getNext())
            previous.setNext(temp)
    def traverse(self):
        current=self.head
        while (current!=None):
            print(current.getData(),end=" ")
            current=current.getNext()
        print()
        
if __name__=='__main__':
    oll=OrderedList()
    oll.add(10)
    oll.add(200)
    oll.add(30)
    print("size",oll.size())
    oll.traverse()
    oll.add(3)
    oll.add(300)
    oll.traverse()
    print("search",oll.search(2000))