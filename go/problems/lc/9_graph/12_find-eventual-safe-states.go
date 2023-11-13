// https://leetcode.com/problems/find-eventual-safe-states/
// https://www.youtube.com/watch?v=uRbJ1OF9aYM&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=20

/*
To find the eventually safe nodes in a graph, we can use a depth-first search (DFS) approach.

1. Create an array to track the visited nodes and initialize all values to false.
2. Create an array to track the current path and initialize all values to false.
3. Iterate through each node in the graph.
4. For each node, check if it has neighbors and if it has not been visited before.
5. If the conditions are met, call the DFS function on that node.
6. In the DFS function, mark the current node as visited and add it to the current path.
7. Recursively call the DFS function on each neighbor of the current node.
8. If a neighbor is already in the current path, it means there is a cycle and the current node is not eventually safe. Mark it in the current path array.
9. After exploring all neighbors, remove the current node from the current path.
10. Once the DFS function returns, check if the current node is eventually safe (i.e., it has been visited but is not in the current path).
11. If the current node is eventually safe, add it to the list of eventually safe nodes.
12. Finally, return the list of eventually safe nodes.

This logic will help identify the nodes in the graph that are eventually safe, 
meaning they do not participate in any cycles in the graph.
*/
package main

func eventualSafeNodes(graph [][]int) []int {
	numNodes := len(graph)
	visitedNodes := make([]bool, numNodes) // Track visited nodes
	currentPath := make([]bool, numNodes)  // Track current path

	// Iterate through each node
	for node, neighborList := range graph {
		// Check if node has neighbors and hasn't been visited before
		if len(neighborList) > 0 && !visitedNodes[node] {
			// Perform depth-first search
			dfs(node, graph, &visitedNodes, &currentPath)
		}
	}
	// Create a list to store the eventually safe nodes
	eventuallySafeNodes := make([]int, 0, numNodes)
	for i := 0; i < numNodes; i++ {
		// Check if the node is visited and in the current path
		if visitedNodes[i] && currentPath[i] {
			continue
		}
		eventuallySafeNodes = append(eventuallySafeNodes, i)
	}
	return eventuallySafeNodes
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
