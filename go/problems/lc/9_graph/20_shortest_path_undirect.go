// https://www.geeksforgeeks.org/problems/shortest-path-in-undirected-graph-having-unit-distance/1

package main

import (
	"fmt"
	"math"
)

func shortestPath(edges [][]int, n, m, src int) []int {
	// Create adjacency matrix
	adj := make([][]bool, n)
	for i := range adj {
		adj[i] = make([]bool, n)
	}

	// Add edges to adjacency matrix
	// Iterate over each edge
	for _, edge := range edges {
		// Check if the two vertices of the edge are not the same i.e. cycle
		if edge[0] != edge[1] {
			// Add the edge to the adjacency matrix in both directions
			adj[edge[0]][edge[1]] = true
			adj[edge[1]][edge[0]] = true
		}
	}

	// Initialize distances array with maximum value
	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt
	}

	visitedNodes := make([]bool, n)
	q := []int{}
	q = append(q, src)
	distances[src] = 0

	// Perform breadth-first search
	for len(q) != 0 {
		node := q[0]
		visitedNodes[node] = true

		// Visit neighbors of the current node
		for neighbor := range adj[node] {
			// Check if the neighbor node has not been visited and there is an edge between the current node and the neighbor
			if !visitedNodes[neighbor] && adj[node][neighbor] {
				// Calculate the distance from the start node to the neighbor
				distance := distances[node] + 1
				// Update the distance if it is shorter than the current distance
				if distances[neighbor] > distance {
					distances[neighbor] = distance
				}
				// Add the neighbor to the queue
				q = append(q, neighbor)
			}
		}

		q = q[1:] // Remove the current node from the queue
	}

	// Convert unreachable nodes to -1
	for k := range distances {
		if distances[k] == math.MaxInt {
			distances[k] = -1
		}
	}

	return distances
}

func main() {
	fmt.Println(shortestPath([][]int{
		{0, 1}, {0, 3}, {3, 4}, {4, 5}, {5, 6}, {1, 2}, {2, 6}, {6, 7}, {7, 8}, {6, 8},
	}, 9, 10, 0))
	fmt.Println(shortestPath([][]int{
		{1, 5},
		{1, 6},
		{2, 0},
		{3, 3},
		{4, 0},
		{4, 6},
		{5, 3},
		{5, 4},
		{6, 5},
		{6, 6},
	}, 7, 10, 4))
}
