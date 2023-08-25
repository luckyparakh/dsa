package main

import "fmt"

type Node struct {
	data  int
	right *Node
	left  *Node
}
type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) Insert(data int) {
	if b.root == nil {
		b.root = &Node{data: data}
	} else {
		b.root.Insert(data)
	}
}

func (n *Node) Insert(data int) {
	if n == nil {
		return
	}
	node := &Node{data: data}
	if n.data >= data {
		if n.left == nil {
			n.left = node
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = node
		} else {
			n.right.Insert(data)
		}
	}
}
func printPreOrder(node *Node) {
	if node == nil {
		return
	}
	nodes := []*Node{}
	nodes = append(nodes, node)
	for len(nodes) != 0 {
		rootNode := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		fmt.Print(rootNode.data, " ")
		if rootNode.right != nil {
			nodes = append(nodes, rootNode.right)
		}
		if rootNode.left != nil {
			nodes = append(nodes, rootNode.left)
		}
	}
	fmt.Println("")
}
func print(ns int, ch rune, node *Node) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", ch, node.data)
	print(ns+2, 'L', node.left)
	print(ns+2, 'R', node.right)
}
func main() {
	bt := BinaryTree{}
	bt.Insert(50)
	bt.Insert(30)
	bt.Insert(-50)
	bt.Insert(60)
	bt.Insert(-60)
	bt.Insert(20)
	bt.Insert(-10)
	bt.Insert(65)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(-5)
	bt.Insert(62)
	bt.Insert(-55)
	bt.Insert(-70)

	// print(0, 'S', bt.root)

	printPreOrder(bt.root)
}
