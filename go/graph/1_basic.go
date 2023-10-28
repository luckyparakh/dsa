package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

type Graph struct {
	Nodes []*Node
}

// addNode adds a new node to the graph if it doesn't already exist
func (g *Graph) addNode(k int) {
	if !verifyVertex(g.Nodes, k) {
		g.Nodes = append(g.Nodes, &Node{Val: k})
	}
}

// getVertex returns the node with the given value if it exists in the graph
func (g *Graph) getVertex(k int) *Node {
	for _, v := range g.Nodes {
		if v.Val == k {
			return v
		}
	}
	return nil
}

// addEdge adds an edge between two nodes in the graph
func (g *Graph) addEdge(from, to int) {
	fromNode := g.getVertex(from)
	toNode := g.getVertex(to)
	if fromNode == nil || toNode == nil {
		return
	}
	if !verifyVertex(fromNode.Neighbors, to) && !verifyVertex(toNode.Neighbors, from) {
		fromNode.Neighbors = append(fromNode.Neighbors, toNode)
		toNode.Neighbors = append(toNode.Neighbors, fromNode)
	}
}

// print prints the graph
func (g *Graph) print() {
	for _, v := range g.Nodes {
		fmt.Printf("\nNode %d: ", v.Val)
		for _, val := range v.Neighbors {
			fmt.Print(val.Val, " ")
		}
	}
	fmt.Println("")
}

// bfs performs a breadth-first search on the graph
func (g *Graph) bfs() {
	if len(g.Nodes) == 0 {
		return
	}
	visitedNode := map[int]bool{}
	visitedNode[g.Nodes[0].Val] = true
	q := []*Node{g.Nodes[0]}
	for len(q) != 0 {
		n := q[0]
		for _, neighbor := range n.Neighbors {
			if !visitedNode[neighbor.Val] {
				visitedNode[neighbor.Val] = true
				q = append(q, neighbor)
			}
		}
		q = q[1:]
		fmt.Println(n.Val)
	}
}

// dfs performs a depth-first search on the graph
func (g *Graph) dfs() {
	if len(g.Nodes) == 0 {
		return
	}
	visitedNode := map[int]bool{}
	visitedNode[g.Nodes[0].Val] = true
	var traverse func(n *Node)
	traverse = func(n *Node) {
		fmt.Println(n.Val)
		for _, neighbor := range n.Neighbors {
			if !visitedNode[neighbor.Val] {
				visitedNode[neighbor.Val] = true
				traverse(neighbor)
			}
		}
	}
	traverse(g.Nodes[0])
}

// verifyVertex checks if a node with the given value exists in the graph
func verifyVertex(g []*Node, k int) bool {
	for _, v := range g {
		if v.Val == k {
			return true
		}
	}
	return false
}

func main() {
	g := Graph{}
	for i := 1; i <= 8; i++ {
		g.addNode(i)
	}
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 5)
	g.addEdge(2, 6)
	g.addEdge(3, 4)
	g.addEdge(3, 7)
	g.addEdge(4, 8)
	g.addEdge(7, 8)
	g.print()
	g.bfs()
	g.dfs()
}