// https://leetcode.com/problems/copy-list-with-random-pointer/
// https://www.youtube.com/watch?v=VNf6VynfpdM&list=PLgUwDviBIf0r47RKH7fdWN54AbWFgGuii&index=15

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
type LinkedList struct {
	head *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	ll := LinkedList{}
	node := head
	var prev *Node
	nodeMap := make(map[*Node]*Node)
	for node != nil {
		newNode := Node{
			Val: node.Val,
		}
		nodeMap[node] = &newNode
		if ll.head == nil {
			ll.head = &newNode
			prev = ll.head
		} else {
			prev.Next = &newNode
			prev = &newNode
		}
		node = node.Next
	}

	node = head
	node1 := ll.head
	for node != nil {
		if node.Random != nil {
			node1.Random = nodeMap[node.Random]
		} else {
			node1.Random = nil
		}
		node = node.Next
		node1 = node1.Next
	}
	fmt.Println(nodeMap)
	n := ll.head
	for n != nil {
		fmt.Print(n.Val, " ")
		n = n.Next
	}
	fmt.Println("")
	return ll.head
}
func (l *LinkedList) Traverse() {
	node := l.head
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println("")
}
func (l *LinkedList) InsertAtEnd(v int) {
	n := Node{
		Val: v,
	}
	if l.head == nil {
		n.Next = nil
		l.head = &n
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &n
	n.Next = nil
}
func copyRandomListWithNoMap(head *Node) *Node {
	// edge
	if head == nil {
		return head
	}

	node := head
	// Insert new node after each node
	for node != nil {
		newNode := Node{
			Val:  node.Val,
			Next: node.Next,
		}
		node.Next = &newNode
		node = newNode.Next
	}

	// Point random pointer of new nodes
	node = head
	for node != nil {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
		node = node.Next.Next
	}

	// Dissect Old and New LL
	newHead := head.Next
	node = head
	for node != nil && node.Next.Next != nil {
		tmp := node.Next.Next
		node.Next.Next = node.Next.Next.Next
		node.Next = tmp
		node = node.Next
	}
	node.Next = nil
	// nh := newHead
	// for nh != nil {
	// 	fmt.Print(" ",nh.Val)
	// 	nh = nh.Next
	// }
	// fmt.Println("")
	return newHead
}
func main() {
	l1 := LinkedList{}
	l1.InsertAtEnd(1)
	l1.InsertAtEnd(10)
	l1.InsertAtEnd(100)
	l1.InsertAtEnd(1000)
	l1.InsertAtEnd(10000)
	l1.Traverse()
	copyRandomListWithNoMap(l1.head)
	l1.Traverse()
}
