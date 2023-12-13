package main

import (
	"fmt"
)

// Function to find number of strongly connected components in the graph
func kosaraju(V int, adj [][]int) int {
	edges := make([][]int, V)
	edgesR := make([][]int, V)
	for i := 0; i < len(adj); i++ {
		edges[adj[i][0]] = append(edges[adj[i][0]], adj[i][1])
		edgesR[adj[i][1]] = append(edgesR[adj[i][1]], adj[i][0])
	}
	fmt.Println(edges)
	fmt.Println(edgesR)
	completions := make([]int, 0)
	vnDFS := make([]bool, V)
	var dfs func(int)
	dfs = func(i int) {
		vnDFS[i] = true
		for _, v := range edges[i] {
			if !vnDFS[v] {
				dfs(v)
			}
		}
		completions = append(completions, i)
	}
	dfs(0)
	fmt.Println(completions)
	vn := make([]bool, V)
	scg := [][]int{}
	for len(completions) != 0 {
		l := len(completions)
		node := completions[l-1]
		if !vn[node] {

			nodes := []int{}
			dfs = func(i int) {
				nodes = append(nodes, i)
				vn[i] = true
				for _, v := range edgesR[i] {
					if !vn[v] {
						dfs(v)
					}
				}
			}
			dfs(node)
			scg = append(scg, nodes)
		}
		completions = completions[:l-1]
	}
	fmt.Println(scg)
	return len(scg)
}
func main() {
	fmt.Println(kosaraju(5, [][]int{{1, 0}, {0, 2}, {2, 1}, {0, 3}, {3, 4}}))
}
