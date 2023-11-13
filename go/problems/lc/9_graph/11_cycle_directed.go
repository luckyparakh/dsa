// https://practice.geeksforgeeks.org/batch/dsa-4/track/DSASP-Graph/problem/detect-cycle-in-a-directed-graph
// https://www.codeconvert.ai/golang-to-java-converter
// https://www.youtube.com/watch?v=9twcmtQj4DU&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=19
package main

// isCyclic checks if a directed graph contains a cycle
func isCyclic(V int, adj [][]int) bool {
	visitedNodes := make([]bool, V) // Track visited nodes
	currentPath := make([]bool, V) // Track current path

	// Iterate through each node
	for node, neighborList := range adj {
		// Check if node has neighbors and hasn't been visited before
		if len(neighborList) > 0 && !visitedNodes[node] {
			// Check if there is a cycle starting from this node
			if dfs(node, adj, &visitedNodes, &currentPath) {
				return true // Cycle found
			}
		}
	}

	return false // No cycle found
}
/*
This code defines a depth-first search function called "dfs" that checks for cycles in a directed graph. 
It takes in a node, an adjacency matrix, and two boolean slices representing visited nodes and the current path. 
The function recursively traverses the graph, marking nodes as visited and checking for cycles. 
If a cycle is found, the function returns true. 
Otherwise, it backtracks and marks the current node as not part of the current path.
*/
func dfs(node int, adj [][]int, visitedNodes, path *[]bool) bool {
	// Get references to the visitedNodes and path slices
	v := *visitedNodes
	p := *path

	// Get the neighbors of the current node
	neighbors := adj[node]

	// Check if the current node has already been visited and is in the current path
	if v[node] && p[node] {
		return true
	}

	// Check if the current node has no neighbors or has already been visited
	if len(neighbors) == 0 || v[node] {
		return false
	}

	// Mark the current node as visited and in the current path
	v[node] = true
	p[node] = true

	// Visit all the neighbors of the current node
	for _, neighbor := range adj[node] {
		if dfs(neighbor, adj, visitedNodes, path) {
			return true
		}
	}

	// Mark the current node as not in the current path
	p[node] = false

	return false
}
