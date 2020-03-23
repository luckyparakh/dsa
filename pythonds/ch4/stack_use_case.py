from ch4.stack import Stack

class stack_use_case(Stack):
    def revstring(self,mystr):
        for item in mystr:
            self.push(item)
        rev_str=""
        while(not self.isEmpty()):
            rev_str=rev_str+self.pop()
        return rev_str
    def bracket_balance(self,expr):
        #st=Stack()
        for item in expr:
            if (item=="("):
                self.push(item)
            elif (item==")"):
                if(not self.isEmpty()):
                    self.pop()
                else:
                    return False
        if(self.isEmpty()):
            return True
        else:
            return False
    def bracket_balance_advance(self,expr):
        open_brackets="([{"
        close_brackets=")]}"
        for item in expr:
            if (item in open_brackets):
                self.push(item)
            elif (item in close_brackets):
                if(not self.isEmpty()):
                    open_bracket=open_brackets[close_brackets.index(item)]
                    poped_item=self.pop()
                    if(poped_item!=open_bracket):
                        return False
                else:
                    return False
        if(self.isEmpty()):
            return True
        else:
            return False
    def baseConvertor(number,base):
        obj=Stack()
        remainder_letter="abcdef"
        while(number!=0):
            rem=number%base
            # in case remainder is greater than 9 then 10 is represented as a & so on.. 
            if(rem>=10): 
                rem=remainder_letter[rem%10]
            obj.push(rem)
            number=number//base
        converted=""
        while(not obj.isEmpty()):
            converted=converted+str(obj.pop())
        return converted
                    
s1=stack_use_case()
print(s1.revstring("abcd"))
s2=stack_use_case()
print(s2.bracket_balance("()()(())"))
s3=stack_use_case()
print(s3.bracket_balance_advance('[{()]'))
print(stack_use_case.baseConvertor(29,16))