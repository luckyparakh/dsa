# -*- coding: utf-8 -*-

from stack import *

def infixToPostfix(expr):
    '''
    Algo
    Scan the token list from left to right.
    If the token is an operand, append it to the end of the output list.
    If the token is a left parenthesis, push it on the opstack.
    If the token is a right parenthesis, pop the opstack until the corresponding left parenthesis is removed. Append each operator to the end of the output list.
    If the token is an operator, *, /, +, or -, push it on the opstack.
    However, first remove any operators already on the opstack that have higher or equal precedence and append them to the output list.
    '''
    list_operator=["**","/","*","+","-"]
    output=[]
    precedence={}
    stk=Stack()
    precedence["**"]=4
    precedence["*"]=3
    precedence["/"]=3
    precedence["+"]=2
    precedence["-"]=2
    precedence["("]=1
    list_token=expr.split()
    for token in list_token:
        if token in list(map(chr,range(65,91))) or token in "0123456789":
            output.append(token)
        elif (token==")"):
            poped_item=stk.pop()
            while (poped_item != "("):
                output.append(poped_item)
                poped_item=stk.pop()
                
        elif (token=="("):
            stk.push(token)
        elif token in list_operator:
            while True:
                if (not stk.isEmpty()): 
                    if (precedence[stk.peek()] >= precedence[token]):
                        output.append(stk.pop())
                    else:
                        stk.push(token)
                        break
                else:
                    stk.push(token)
                    break
    while (not stk.isEmpty()):
        output.append(stk.pop())
    return ' '.join(output)
def calculation(operator,op1,op2):
    if operator in "+":
        return op2+op1
    elif operator=="-":
        return op2-op1
    elif operator=="*":
        return op2-op1
    elif operator=="/":
        return op2/op1
    
    
    
def PostfixEval(expr):
    '''
    If the token is an operand, convert it from a string to an integer and push the value onto the operandStack.
    If the token is an operator, *, /, +, or -, it will need two operands. Pop the operandStack twice.
    The first pop is the second operand and the second pop is the first operand. Perform the arithmetic operation.
    Push the result back on the operandStack.
    '''
    #print(expr)
    list_operator=["/","*","+","-"]
    stk=Stack()
    for token in expr:
        if token in "0123456789":
            stk.push(int(token))
        elif token in list_operator:
            op1=stk.pop()
            op2=stk.pop()
            stk.push(calculation(token,op1,op2))
    return stk.pop()
        
print(infixToPostfix("A * B + C * D"))
print(infixToPostfix("( A + B ) * C - ( D - E ) * ( F + G )"))
print(infixToPostfix("5 * 3 ** ( 4 - 2 )"))
print(PostfixEval("5 3 4 2 - + *"))
print(PostfixEval('7 8 + 3 2 + /'))
