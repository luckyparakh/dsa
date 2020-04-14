
#0,1,1,2,3,5,8
def getNthFib(n):
    fib_dict={1:0,2:1}
    if n in fib_dict.keys():
        op=fib_dict[n]
    else:
        for i in range(3,n+1):
            op=fib_dict[i-1]+fib_dict[i-2]
            fib_dict[i]=op
    return op
            
print(getNthFib(6))

