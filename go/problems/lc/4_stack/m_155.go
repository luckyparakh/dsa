package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	items []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		newVal := 2*val - this.min
		this.min = val
		this.items = append(this.items, newVal)
	} else {
		this.items = append(this.items, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.items) == 0 {
		return
	}
	l := len(this.items) - 1
	if this.items[l] < this.min {
		this.min = 2*this.min - this.items[l]
	}
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	topVal := this.items[len(this.items)-1]
	if topVal < this.min {
		return this.min
	}
	return topVal
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	ms := Constructor()
	ms.Push(5)
	ms.Push(3)
	// ms.Push(-3)
	fmt.Println(ms)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms)
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())
}
