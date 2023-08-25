package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type LinkedList struct {
	head *Node
	Tail *Node
}

func (this *LinkedList) Delete() {
	if this.Tail == nil {
		return
	}
	if this.Tail.Prev != nil {
		this.Tail = this.Tail.Prev
		this.Tail.Next = nil
	} else {
		this.Tail = nil
		this.head = nil
	}
}
func (this *LinkedList) Insert(v int) {
	n := &Node{Val: v}
	if this.head == nil {
		this.Tail = n
		this.head = n
	} else {
		this.head.Prev = n
		n.Next = this.head
		n.Prev = nil
		this.head = n
	}
}
func (l *LinkedList) Traverse() {
	if l.head == nil {
		fmt.Println("Empty list")
		return
	}
	current := l.head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
func main() {
	dll := LinkedList{}
	dll.Traverse()
	dll.Insert(5)
	dll.Insert(4)
	dll.Insert(3)
	dll.Traverse()
	dll.Delete()
	dll.Traverse()
	dll.Delete()
	dll.Traverse()
	dll.Delete()
	dll.Traverse()
}
