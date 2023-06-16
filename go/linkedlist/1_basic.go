package main

import "fmt"

// Linked list are best for insertion and deletion of items.
// Traverse a linked list takes o(n)

// Always keep scenario of empty list in mind while traversing, inserting or deleting items
// Always keep scenario of list with one element in mind while traversing, inserting or deleting items

type node struct {
	next *node
	data string
}

type ll struct {
	length int
	head   *node
}

func initLinkList() *ll {
	return &ll{}
}

func (l *ll) AddFront(d string) {
	element := &node{
		data: d,
		next: l.head,
	}
	l.head = element
	l.length++
}

func (l *ll) RemoveFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	l.length--
}

func (l *ll) AddBack(d string) {
	element := &node{
		data: d,
	}
	if l.head == nil {
		element.next = l.head
		l.head = element
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		element.next = current.next
		current.next = element
	}
	l.length++
}

func (l *ll) RemoveBack() {
	if l.head == nil {
		return
	}
	current := l.head
	var previous *node
	for current.next != nil {
		previous = current
		current = current.next
	}
	if previous != nil {
		previous.next = nil
	} else {
		// if ll has only one element before removal of item
		l.head = nil
	}
	l.length--
}

func (l *ll) Traverse() {
	if l.head == nil {
		fmt.Println("Empty list")
		return
	}
	current := l.head
	for current != nil {
		fmt.Print(current.data + " ")
		current = current.next
	}
	fmt.Println()
}

func (l *ll) Len() {
	fmt.Println(l.length)
}

func main() {
	l := initLinkList()

	l.Traverse()
	l.AddFront("b")
	l.AddFront("a")
	l.AddBack("z")
	l.Traverse()
	l.Len()
	l.RemoveBack()
	l.RemoveFront()
	l.Traverse()
	l.Len()
	l.RemoveBack()
	l.Traverse()
	l.Len()
	l.RemoveBack()
	l.Traverse()
	l.Len()
	l.RemoveFront()
	l.Traverse()
	l.Len()
}
