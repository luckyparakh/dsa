// https://www.geeksforgeeks.org/problems/shortest-path-in-undirected-graph/1
// https://www.youtube.com/watch?v=ZUFQfFaU-8U&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=27
// https://thepythoncode.com/assistant/code-converter/python/
package main

import (
	"fmt"
	"math"
)

func shortestPath(n int, m int, edges [][]int) []int {
	// Create adjacency matrix
	adj := make([][][]int, n)
	for i := range adj {
		adj[i] = make([][]int, 0)
	}
	for _, edge := range edges {
		// Add edge to adjacency matrix
		adj[edge[0]] = append(adj[edge[0]], []int{edge[1], edge[2]})
	}

	// Initialize distances array with maximum value
	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt
	}

	// Perform topological sort
	ts := topoSort(adj)

	// Set distance from source node to 0
	distances[0] = 0

	// Calculate shortest distances
	for len(ts) != 0 {
		node := ts[0]
		for _, neighbor := range adj[node] {
			distance := distances[node] + neighbor[1]
			if distances[neighbor[0]] > distance {
				// Update distance if shorter path is found
				distances[neighbor[0]] = distance
			}
		}
		ts = ts[1:]
	}

	// Replace maximum distance with -1
	for i := range distances {
		if distances[i] == math.MaxInt {
			distances[i] = -1
		}
	}

	return distances
}

func topoSort(adj [][][]int) []int {
	// nodes is the number of nodes in the graph
	nodes := len(adj)
	// visitedNodes keeps track of which nodes have been visited during the DFS
	visitedNodes := make([]bool, nodes)
	// temp holds the nodes in the order they were visited during the DFS
	temp := []int{}
	var dfs func(k int)
	dfs = func(k int) {
		visitedNodes[k] = true
		for _, neighbor := range adj[k] {
			if !visitedNodes[neighbor[0]] {
				dfs(neighbor[0])
			}
		}
		temp = append(temp, k)
	}
	for k := range adj {
		if !visitedNodes[k] {
			dfs(k)
		}
	}
	// fmt.Println(temp)
	ans := []int{}
	for len(temp) != 0 {
		lastNodeIndex := len(temp) - 1
		ans = append(ans, temp[lastNodeIndex])
		temp = temp[:lastNodeIndex]
	}
	return ans
}

func main() {
	output := shortestPath(4, 2, [][]int{
		{0, 1, 2}, {0, 2, 1},
	})
	fmt.Println(output) //[0,2,1,-1]
}
