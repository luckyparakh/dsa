package main

type Node struct {
	Val  int
	Next *Node
	Random *Node
}



// Create Map of nodes
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodesMap := make(map[*Node]*Node)
	n := head
	for n != nil {
		nodesMap[n] = &Node{Val: n.Val, Random: nil, Next: nil}
		n = n.Next
	}
	n = head
	for n != nil {
		newNode := nodesMap[n]
		newNode.Random = nodesMap[n.Random]
		newNode.Next = nodesMap[n.Next]
		n = n.Next
	}
	return nodesMap[head]
}
