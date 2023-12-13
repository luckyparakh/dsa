// https://www.geeksforgeeks.org/problems/minimum-spanning-tree/1

package main

import (
	"fmt"
	"sort"
)

type Node struct {
	number int
	parent int
	weight int
}

func spanningTree(V int, adj [][]int) int {
	// Initialize an empty graph
	graph := make([][][]int, V)
	for i := 0; i < V; i++ {
		graph[i] = make([][]int, 0)
	}
	// Add all edges to the graph
	for i := 0; i < len(adj); i++ {
		u, v, w := adj[i][0], adj[i][1], adj[i][2]
		graph[u] = append(graph[u], []int{v, w})
		graph[v] = append(graph[v], []int{u, w})
	}
	// fmt.Println(graph)
	// Initialize a visited array
	visited := make([]bool, V)

	// Initialize a queue
	q := make([]Node, 0)
	q = append(q, Node{0, -1, 0})
	mst := [][]int{}
	sum := 0
	for len(q) != 0 {
		sort.SliceStable(q, func(i, j int) bool {
			return q[i].weight < q[j].weight
		})
		node := q[0]
		q = q[1:]
		if !visited[node.number] {
			for _, n := range graph[node.number] {
				q = append(q, Node{n[0], node.number, n[1]})
			}
			// fmt.Println(q)
			visited[node.number] = true
			sum += node.weight
			// fmt.Println(sum)
			if node.parent != -1 {
				mst = append(mst, []int{node.parent, node.number})
			}
		}
	}
	// fmt.Println(mst)
	return sum
}

func main() {
	fmt.Println(spanningTree(3, [][]int{{0, 1, 5}, {1, 2, 3}, {0, 2, 1}})) //4
	fmt.Println(spanningTree(2, [][]int{{0, 1, 5}})) //5
}
