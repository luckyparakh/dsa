// https://www.youtube.com/watch?v=vXrv3kruvwE
// https://practice.geeksforgeeks.org/batch/dsa-4/track/DSASP-Graph/problem/detect-cycle-in-an-undirected-graph
// https://www.codeconvert.ai/golang-to-java-converter

// https://www.youtube.com/watch?v=vXrv3kruvwE
// https://practice.geeksforgeeks.org/batch/dsa-4/track/DSASP-Graph/problem/detect-cycle-in-an-undirected-graph
// https://www.codeconvert.ai/golang-to-java-converter

package main

import "fmt"

// isCycle checks if there is a cycle in the given graph
func isCycle(V int, adj [][]int) bool {
	// Three state 0 not visited, -1 visited, 1 processed
	visitedNodes := make([]int, V) // Track visited nodes

	// Iterate through each node in the graph
	for node, neighborList := range adj {
		// If the node has neighbors and has not been visited yet
		if len(neighborList) > 0 && visitedNodes[node] == 0 {
			// Check if there is a cycle starting from this node
			if bfs(node, adj, &visitedNodes) {
				return true
			}
		}
	}

	return false
}

// bfs performs a breadth-first search to check for cycles
func bfs(node int, adj [][]int, visitedNodes *[]int) bool {
	queue := []int{}
	queue = append(queue, node)
	visited := (*visitedNodes)

	// Process nodes in the queue
	for len(queue) != 0 {
		currentNode := queue[0]

		// Check each neighbor of the current node
		for _, neighbor := range adj[currentNode] {
			// If the neighbor has already been visited, there is a cycle
			if visited[neighbor] == -1 {
				return true
			}
			// If the neighbor has not been visited yet, add it to the queue and mark it as visited
			if visited[neighbor] == 0 {
				queue = append(queue, neighbor)
				visited[neighbor] = -1
			}
		}

		// Mark the current node as processed
		visited[currentNode] = 1

		// Remove the current node from the queue
		queue = queue[1:]
	}

	return false
}

func main() {
	// Test the isCycle function
	fmt.Println(isCycle(5, [][]int{
		{1}, {0, 2, 4}, {1, 3}, {2, 4}, {1, 3},
	}))
}