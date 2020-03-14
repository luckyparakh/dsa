# -*- coding: utf-8 -*-
import random
class die:
    def __init__(self,num_sides):
        self.num_sides=num_sides
        self.current_value=self.roll()
    def roll(self):
        self.current_value=random.randrange(1,self.num_sides+1)
        return self.current_value
    def __str__(self):
        return str(self.current_value)
    def __repr__(self):
        return "Die ({}):{}".format(self.num_sides,self.current_value)
    def __eq__(self,other):
        return self.current_value == other.current_value

my_die=die(5)
for i in range(5):
    print(my_die,my_die.current_value)
    my_die.roll()
d_list=[die(6),die(20)]
print(d_list)
print(die(6)==die(20))