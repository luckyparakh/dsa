# -*- coding: utf-8 -*-

from ch4.queue import queue

def hot_potato(names,num):
    q=queue()
    i=0
    for name in names:
        q.enqueue(name)
    while (len(q.item)>1):
        if(i%num==0):
            q.dequeue()
        else:
            q.enqueue(q.dequeue())
        i=i+1
    return q.dequeue()

hot_potato(["Bill","David","Susan","Jane","Kent","Brad"],7)