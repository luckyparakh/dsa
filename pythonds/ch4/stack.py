# -*- coding: utf-8 -*-

class Stack:
    def __init__(self):
        self.stk=[]
    def push(self,value):
        self.stk.append(value) # O(1)
    def pop(self):
        return self.stk.pop() # O(1)
    def isEmpty(self):
        if(self.size()==0):
            return True
        else:
            return False
    def size(self):
        return len(self.stk)
    def peek(self):
        if (self.isEmpty()!=True):
            return self.stk[-1]
        else:
            return "Empty"
    
if (__name__=='__main__'):      
    s=Stack()
    print(s.isEmpty())
    print(s.peek())
    s.push(4)
    s.push('dog')
    print(s.peek())
    s.push(True)
    print(s.isEmpty())
    s.push(8.4)
    print(s.pop())
    print(s.size())

